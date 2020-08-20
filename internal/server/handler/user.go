package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/lifina/go-template/internal/server/service"
	"go.uber.org/zap"
)

type UserHandler struct {
	srv    service.User
	logger *zap.Logger
}

type userCraeteParameter struct {
	Name string `json:"name"`
	Age  uint8  `json:"age"`
}

func (u *UserHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "user_id")
	user, err := u.srv.GetByID(r.Context(), id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		if _, err := w.Write([]byte("error")); err != nil {
			u.logger.Error("handler: write to writer failed", zap.Error(err))
		}
		return
	}
	b, err := json.Marshal(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		if _, err := w.Write([]byte("error")); err != nil {
			u.logger.Error("handler: write to writer failed", zap.Error(err))
		}
		return
	}
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(b); err != nil {
		u.logger.Error("handler: write to writer failed", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (u *UserHandler) Create(w http.ResponseWriter, r *http.Request) {

	param := userCraeteParameter{}
	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(&param); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		if _, err := w.Write([]byte("error")); err != nil {
			u.logger.Error("handler: write to writer failed", zap.Error(err))
			w.WriteHeader(http.StatusBadRequest)
		}
		return
	}
	if err := u.srv.Create(r.Context(), param.Name, param.Age); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		if _, err := w.Write([]byte("error")); err != nil {
			u.logger.Error("handler: write to writer failed", zap.Error(err))
		}
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func NewUserHandler(srv service.User, logger *zap.Logger) *UserHandler {
	return &UserHandler{
		srv:    srv,
		logger: logger,
	}
}
