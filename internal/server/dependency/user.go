package dependency

import (
	"database/sql"

	"github.com/go-chi/chi"
	"github.com/lifina/go-template/internal/pkg/id"
	"github.com/lifina/go-template/internal/server/handler"
	"github.com/lifina/go-template/internal/server/infrastructure/repository"
	"github.com/lifina/go-template/internal/server/routes"
	"github.com/lifina/go-template/internal/server/service"
	"go.uber.org/zap"
)

func NewUser(r *chi.Mux, db *sql.DB, logger *zap.Logger) {
	idgen := id.NewUUIDV4Generator()
	repo := repository.NewUserReposiotry(db, logger, idgen)
	service := service.NewUserService(repo, logger)
	userHandler := handler.NewUserHandler(service, logger)
	u := routes.User{Get: userHandler.GetByID, Post: userHandler.Create}
	u.Install(r)
}
