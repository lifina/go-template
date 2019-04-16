package service

import (
	"context"

	"github.com/lifina/go-template/pkg/model"
	"go.uber.org/zap"
)

type User interface {
	Get(ctx context.Context, id int64) (*model.User, error)
	Regist(ctx context.Context, name, sex string, age int64) error
}

type userService struct {
	repo   model.UserRepository
	logger *zap.Logger
}

func NewUserService(repo model.UserRepository, logger *zap.Logger) User {
	return &userService{
		repo:   repo,
		logger: logger,
	}
}

func (u *userService) Get(ctx context.Context, id int64) (*model.User, error) {
	return u.repo.Get(ctx, id)
}

func (u *userService) Regist(ctx context.Context, name, sex string, age int64) error {
	m := &model.User{
		Name: name,
		Sex:  sex,
		Age:  age,
	}
	return u.repo.Create(ctx, m)
}
