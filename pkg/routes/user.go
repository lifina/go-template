package routes

import (
	"net/http"

	"github.com/go-chi/chi"
)

type User struct {
	Get  http.HandlerFunc
	Post http.HandlerFunc
}

func (u *User) Install(r *chi.Mux) {
	r.Post("/user", u.Post)
	r.Get("/user/{user_id}", u.Get)
}
