package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type DocumentRoutes struct{}

func (rs DocumentRoutes) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world from document"))
	})

	return r
}
