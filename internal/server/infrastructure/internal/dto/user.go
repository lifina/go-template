package dto

import (
	"time"

	"github.com/lifina/go-template/internal/server/domain/user"
)

type User struct {
	ID        string
	Name      string
	Age       uint8
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (e *User) OutputModel() *user.User {
	return &user.User{
		ID:   user.ID(e.ID),
		Name: e.Name,
		Age:  e.Age,
	}
}
