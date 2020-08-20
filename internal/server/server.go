package server

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-sql-driver/mysql"
	"github.com/lifina/go-template/internal/pkg/logger"
	"github.com/lifina/go-template/internal/pkg/utils"
	"github.com/lifina/go-template/internal/server/dependency"
	"github.com/lifina/go-template/internal/server/handler"
	"github.com/lifina/go-template/internal/server/routes"
	"go.uber.org/zap"
)

type Server struct {
	*http.Server
}

type Config struct {
	AppName    string
	Revision   string
	Port       string
	DBHost     string
	DBPort     int64
	DBNet      string
	DBUser     string
	DBPassword string
	DBName     string
}

func NewServer(config *Config) (*Server, error) {

	logger, err := logger.NewLogger(config.AppName, config.Revision)
	if err != nil {
		return nil, err
	}

	mux := chi.NewMux()
	indexRoute := &routes.Index{IndexHandler: handler.NewPingHandler(logger).Ping}
	indexRoute.Install(mux)

	db, err := sql.Open("mysql", config.OutputDBConfig().FormatDSN())
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		logger.Error(
			"server: connection to database failed",
			zap.String("dsn", config.OutputDBConfig().FormatDSN()),
			zap.Error(err),
		)
		return nil, err
	}

	dependency.NewUser(mux, db, logger)

	server := http.Server{
		Addr:    fmt.Sprintf(":%s", config.Port),
		Handler: mux,
	}

	return &Server{&server}, nil
}

func (c *Config) OutputDBConfig() *mysql.Config {
	return &mysql.Config{
		User:                 c.DBUser,
		Passwd:               c.DBPassword,
		Addr:                 c.DBHost,
		Net:                  c.DBNet,
		DBName:               c.DBName,
		Loc:                  utils.TimeZoneJST,
		AllowNativePasswords: true,
		ParseTime:            true,
	}
}
