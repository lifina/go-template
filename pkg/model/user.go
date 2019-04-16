package model

import "context"

type User struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Age  int64  `json:"age"`
	Sex  string `json:"sex"`
}

func (u *User) IsAdult() bool {
	return u.Age >= 18
}

type UserRepository interface {
	Get(ctx context.Context, id int64) (*User, error)
	Create(ctx context.Context, user *User) error
}
