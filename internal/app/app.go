package app

import (
	"auth-service/config"
	handlers "auth-service/internal/handlers/http/v1"
	"auth-service/internal/repository"
	"auth-service/internal/service"
	"context"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Run() error {
	var logger *zap.Logger
	var err error
	appMode := os.Getenv("APP_MODE")
	if appMode == "prod" {
		logger, err = zap.NewProduction()
		if err != nil {
			log.Println("Error while creating logger: ", err)
			return err
		}
	} else if appMode == "dev" {
		logger, err = zap.NewDevelopment()
		if err != nil {
			log.Println("Error while creating logger: ", err)
			return err
		}
	} else {
		log.Println("Error while creating logger: logger mode undefined")
		return fmt.Errorf("error while defining logger: app_mode is invalid %s", appMode)
	}

	defer logger.Sync()
	sugar := logger.Sugar()
	cfg, err := config.New()
	if err != nil {
		sugar.Errorf("errof while defining config: %v", err)
		return err
	}
	db, err := repository.NewPostgresPool(cfg.DB)
	if err != nil {
		sugar.Errorf("error while creating database: %v", err)
		return err
	}
	defer db.Close()

	repos := repository.NewRepository(db, time.Duration(cfg.DB.TimeOut)*time.Second, sugar)
	if err != nil {
		sugar.Errorf("error while creating repository: %v", err)
		return err
	}
	services := service.NewService(repos, cfg, sugar)

	h := handlers.NewHandler(services, sugar, cfg)
	logger.Info(cfg.App.AppPort)
	srv := http.Server{ //nolint:gosec
		Addr:    ":" + cfg.App.AppPort,
		Handler: h.InitRoutes(),
	}

	errChan := make(chan error, 1)

	go func(errChan chan<- error) {
		sugar.Infof("Starting server on port: %s\n", cfg.App.AppPort)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			sugar.Error(err.Error())
			errChan <- err
		}
	}(errChan)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	select {
	case <-quit:
		sugar.Error("Killing signal was received")
	case err := <-errChan:
		sugar.Errorf("HTTP server run error: %s", err)
	}

	sugar.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(cfg.App.AppShutdownTime))
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		sugar.Infof("Server forced to shutdown: %s", err)
	}
	return nil
}
