package api

import (
	"github.com/go-chi/chi"
	"github.com/lifina/go-template/pkg/database"
	"github.com/lifina/go-template/pkg/di"
	"github.com/lifina/go-template/pkg/handler"
	"github.com/lifina/go-template/pkg/logger"
	"github.com/lifina/go-template/pkg/routes"
)

func New(dbAddr string) (*chi.Mux, error) {
	logger := logger.NewLogger()

	// inject
	r := chi.NewRouter()
	i := routes.Index{IndexHandler: handler.Ping}
	i.Install(r)

	db, err := database.NewClient(dbAddr)
	if err != nil {
		return nil, err
	}

	di.NewUser(r, db, logger)

	return r, nil
}
