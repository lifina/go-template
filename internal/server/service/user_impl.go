package service

import (
	"context"

	"github.com/lifina/go-template/internal/server/domain/user"
	"go.uber.org/zap"
)

type userService struct {
	repo   user.Repository
	logger *zap.Logger
}

func (u *userService) GetByID(ctx context.Context, uid string) (*user.User, error) {
	return u.repo.GetByID(ctx, user.ID(uid))
}

func (u *userService) Create(ctx context.Context, name string, age uint8) error {
	uid := u.repo.NextID(ctx)
	m, err := user.NewUser(uid, name, age)
	if err != nil {
		return err
	}
	return u.repo.Create(ctx, m)
}

func NewUserService(repo user.Repository, logger *zap.Logger) User {
	return &userService{
		repo:   repo,
		logger: logger,
	}
}
