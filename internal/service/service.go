package service

import (
	"auth-service/config"
	"auth-service/internal/models"
	"auth-service/internal/repository"
	"go.uber.org/zap"
)

type Service struct {
	User
}

type User interface {
	CreateUser(creds *models.SignUpUser) error
	SignInUser(creds *models.SignInUser) error
	SignOutUser(userPublicID string) error
}

func NewService(repo *repository.Repository, cfg *config.Configs, logger *zap.SugaredLogger) *Service {
	return &Service{
		User: newUserService(repo.User, cfg, logger),
	}
}
