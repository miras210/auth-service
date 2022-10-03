package service

import (
	"auth-service/config"
	"auth-service/internal/repository"
	"go.uber.org/zap"
)

type Service struct {
}

func NewService(repo *repository.Repository, cfg *config.Configs, logger *zap.SugaredLogger) *Service {
	return &Service{}
}
