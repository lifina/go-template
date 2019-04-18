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

type userEntity struct {
	ID   int64
	Name string
	Age  int64
	Sex  string
}

func (u *userEntity) buildFromModel(m model.User) {
	u.ID = m.ID
	u.Name = m.Name
	u.Age = m.Age
	u.Sex = m.Sex
}

func (u *userEntity) generateModel() *model.User {
	return &model.User{
		ID:   u.ID,
		Name: u.Name,
		Age:  u.Age,
		Sex:  u.Sex,
	}
}
