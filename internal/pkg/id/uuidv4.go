package id

import "github.com/google/uuid"

type uuidv4Generator struct {
}

func (g *uuidv4Generator) Generate() string {
	return uuid.New().String()
}

func NewUUIDV4Generator() Generator {
	return &uuidv4Generator{}
}
