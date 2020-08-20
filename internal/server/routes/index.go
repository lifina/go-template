package routes

import (
	"net/http"

	"github.com/go-chi/chi"
)

type Index struct {
	IndexHandler http.HandlerFunc
}

func (i *Index) Install(r *chi.Mux) {
	r.Get("/", i.IndexHandler)
}
