package repository

import (
	"auth-service/internal/models"
	"context"
	"database/sql"
	"go.uber.org/zap"
	"time"
)

type UserRepository struct {
	db      *sql.DB
	timeout time.Duration
	logger  *zap.SugaredLogger
}

func newUserRepository(db *sql.DB, timeout time.Duration, logger *zap.SugaredLogger) *UserRepository {
	return &UserRepository{
		db:      db,
		timeout: timeout,
		logger:  logger,
	}
}

func (u *UserRepository) CreateUser(user *models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), u.timeout)
	defer cancel()

	query := `INSERT INTO users (username, password) VALUES($1, $2) RETURNING id`
	if err := u.db.QueryRowContext(ctx, query, user.Username, user.Password).Scan(&user.ID); err != nil {
		u.logger.Errorf("Error occurred while querying to DB: %s", err.Error())
		return err
	}
	u.logger.Infof("User successfully created with ID: %d", user.ID)
	return nil
}

func (u *UserRepository) GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	ctx, cancel := context.WithTimeout(context.Background(), u.timeout)
	defer cancel()

	query := `SELECT public_id, username, password FROM users WHERE username = $1`
	if err := u.db.QueryRowContext(ctx, query, username).Scan(&user.PublicID, &user.Username, &user.Password); err != nil {
		u.logger.Errorf("Error occurred while querying to DB: %s", models.ErrUserDoesNotExist)
		return nil, err
	}
	u.logger.Infof("User successfully retrieved with ID: %d", user.ID)

	return &user, nil
}
