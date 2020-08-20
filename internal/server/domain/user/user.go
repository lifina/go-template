package user

import (
	"context"
	"errors"
)

type Repository interface {
	GetByID(ctx context.Context, id ID) (*User, error)
	Create(ctx context.Context, user *User) error
	NextID(ctx context.Context) ID
}

var (
	ErrUserNotFound = errors.New("reposiotry: user not found")
)

type User struct {
	ID   ID     `json:"id"`
	Name string `json:"name"`
	Age  uint8  `json:"age"`
}

type ID string

func (u ID) String() string {
	return string(u)
}

func NewUser(id ID, name string, age uint8) (*User, error) {
	return &User{
		ID:   id,
		Name: name,
		Age:  age,
	}, nil
}
