package entity

import "github.com/lifina/go-template/pkg/model"

type User struct {
	ID   int64
	Name string
	Age  int64
	Sex  string
}

func BuildFromModel(u model.User) *User {
	return &User{
		ID:   u.ID,
		Name: u.Name,
		Age:  u.Age,
		Sex:  u.Sex,
	}
}

func (u *User) GenerateModel() *model.User {
	return &model.User{
		ID:   u.ID,
		Name: u.Name,
		Age:  u.Age,
		Sex:  u.Sex,
	}
}
