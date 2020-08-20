package logger

import (
	"github.com/lifina/go-template/internal/pkg/logger/internal/stackdriver"
	"go.uber.org/zap"
)

func NewLogger(name, revision string) (*zap.Logger, error) {
	return stackdriver.NewLogger(name, revision)
}

func NewLoggerDevelopment(name, revision string) (*zap.Logger, error) {
	return stackdriver.NewLoggerDevelopment(name, revision)
}
