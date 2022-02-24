package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type AdminRoutes struct{}

func (rs AdminRoutes) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world from admin"))
	})

	return r
}
