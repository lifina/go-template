package di

import (
	"github.com/go-chi/chi"
	"github.com/jmoiron/sqlx"
	"github.com/lifina/go-template/pkg/handler"
	"github.com/lifina/go-template/pkg/repository"
	"github.com/lifina/go-template/pkg/routes"
	"github.com/lifina/go-template/pkg/service"
	"go.uber.org/zap"
)

func NewUser(r *chi.Mux, db *sqlx.DB, logger *zap.Logger) {
	repo := repository.NewUserRepository(db, logger)
	service := service.NewUserService(repo, logger)
	userHandler := handler.NewUserHandler(service, logger)
	u := routes.User{Get: userHandler.Get, Post: userHandler.Post}
	u.Install(r)
}
