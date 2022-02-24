package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type ConfigurationRoutes struct{}

func (rs ConfigurationRoutes) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world from config"))
	})

	return r
}
