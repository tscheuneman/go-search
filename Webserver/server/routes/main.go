package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type MainRouter struct{}

func (rs MainRouter) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("main root router"))
	})

	r.Mount("/admin", AdminRoutes{}.Routes())

	return r
}
