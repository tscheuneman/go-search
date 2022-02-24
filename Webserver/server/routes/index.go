package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type IndexRoutes struct{}

func (rs IndexRoutes) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world from index"))
	})

	return r
}
