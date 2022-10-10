package service

import (
	"auth-service/config"
	"auth-service/internal/models"
	"auth-service/internal/repository"
	"go.uber.org/zap"
)

type UserService struct {
	userRepo repository.User
	cfg      *config.Configs
	log      *zap.SugaredLogger
}

func newUserService(userRepo repository.User, cfg *config.Configs, log *zap.SugaredLogger) *UserService {
	return &UserService{
		userRepo: userRepo,
		cfg:      cfg,
		log:      log,
	}
}

func (u *UserService) CreateUser(creds *models.SignUpUser) error {
	user := &models.User{
		Username: creds.Username,
		Password: creds.Password,
	}

	err := u.userRepo.CreateUser(user)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserService) SignInUser(creds *models.SignInUser) error {
	//TODO implement me
	panic("implement me")
}

func (u *UserService) SignOutUser(userPublicID string) error {
	//TODO implement me
	panic("implement me")
}
