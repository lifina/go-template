package logger

import (
	"github.com/blendle/zapdriver"
	"github.com/lifina/go-template/pkg/config"
	"go.uber.org/zap"
)

func NewLogger() *zap.Logger {
	logger, _ := zapdriver.NewDevelopment()
	if config.IsEnvProduction() {
		logger, _ = zapdriver.NewProduction()
	}
	return logger
}
