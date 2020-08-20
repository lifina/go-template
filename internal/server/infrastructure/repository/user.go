package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/lifina/go-template/internal/pkg/id"
	"github.com/lifina/go-template/internal/server/domain/user"
	"github.com/lifina/go-template/internal/server/infrastructure/internal/dto"

	"go.uber.org/zap"
)

type userRepository struct {
	db     *sql.DB
	logger *zap.Logger
	idgen  id.Generator
}

func (u *userRepository) GetByID(ctx context.Context, id user.ID) (*user.User, error) {
	query := "select id, name, age from users where id = ?"
	row := u.db.QueryRowContext(ctx, query, id)

	userDTO := &dto.User{}
	err := row.Scan(&userDTO.ID, &userDTO.Name, &userDTO.Age)
	switch err {
	case nil:
		return userDTO.OutputModel(), nil
	case sql.ErrNoRows:
		u.logger.Warn("user: user not found", zap.Error(err))
		return nil, user.ErrUserNotFound
	default:
		u.logger.Error("user: failed to get user", zap.Error(err))
		return nil, err
	}
}

func (u *userRepository) Create(ctx context.Context, userRepository *user.User) error {
	query := "insert into users (`id`, `name`, `age`) values (?, ?, ?)"
	u.logger.Info("DEBUG:", zap.String("query", query), zap.String("user", fmt.Sprintf("%+v", userRepository)))
	_, err := u.db.ExecContext(ctx, query, userRepository.ID, userRepository.Name, userRepository.Age)
	if err != nil {
		u.logger.Error("user: failed to create user", zap.Error(err))
		return err
	}
	return nil
}

func (u *userRepository) NextID(ctx context.Context) user.ID {
	return user.ID(u.idgen.Generate())
}

func NewUserReposiotry(db *sql.DB, logger *zap.Logger, idgen id.Generator) user.Repository {
	return &userRepository{
		db:     db,
		logger: logger,
		idgen:  idgen,
	}
}
