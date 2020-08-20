package service

import (
	"context"

	"github.com/lifina/go-template/internal/server/domain/user"
)

type User interface {
	GetByID(context.Context, string) (*user.User, error)
	Create(context.Context, string, uint8) error
}
