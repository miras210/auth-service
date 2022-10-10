package repository

import (
	"auth-service/internal/models"
	"database/sql"
	"go.uber.org/zap"
	"time"
)

type Repository struct {
	User
}

type User interface {
	CreateUser(user *models.User) error
	GetUserByUsername(username string) (*models.User, error)
}

func NewRepository(db *sql.DB, timeout time.Duration, logger *zap.SugaredLogger) *Repository {
	return &Repository{
		User: newUserRepository(db, timeout, logger),
	}
}
