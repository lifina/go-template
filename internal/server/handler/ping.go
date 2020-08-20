package handler

import (
	"net/http"

	"go.uber.org/zap"
)

type PingHandler struct {
	logger *zap.Logger
}

func (p *PingHandler) Ping(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write([]byte("pongpong")); err != nil {
		p.logger.Error("handler: write to writer failed", zap.Error(err))
	}
}

func NewPingHandler(logger *zap.Logger) *PingHandler {
	return &PingHandler{
		logger: logger,
	}
}
