package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/lifina/go-template/pkg/model"
	"go.uber.org/zap"
)

type user struct {
	db     *sqlx.DB
	logger *zap.Logger
}

func NewUserRepository(db *sqlx.DB, logger *zap.Logger) model.UserRepository {
	return &user{
		db:     db,
		logger: logger,
	}
}

func (u *user) Get(ctx context.Context, id int64) (*model.User, error) {
	return nil, nil
}

func (u *user) Create(ctx context.Context, user *model.User) error {
	return nil
}
