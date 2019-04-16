package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"go.uber.org/zap"

	"github.com/lifina/go-template/pkg/service"
)

type UserHandler struct {
	service service.User
	logger  *zap.Logger
}

func NewUserHandler(svc service.User, logger *zap.Logger) *UserHandler {
	return &UserHandler{
		service: svc,
		logger:  logger,
	}
}

func (u *UserHandler) Get(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "user_id"), 10, 64)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("error"))
	}
	user, err := u.service.Get(r.Context(), id)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("error"))
	}
	b, err := json.Marshal(user)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("error"))
	}
	w.WriteHeader(200)
	w.Write(b)
}

func (u *UserHandler) Post(w http.ResponseWriter, r *http.Request) {

}
