package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type AdminRoutes struct{}

func (rs AdminRoutes) Routes() chi.Router {
	r := chi.NewRouter()

	// TODO: Soooome kind of auth middleware.
	// It'll probably be something super basic, probably no roles/ granular permissions

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world from admin"))
	})

	// Sub routes for indexes actions
	r.Mount("/index", IndexRoutes{}.Routes())

	// Auth eventually anyway...
	r.Get("/login", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Login Func"))
	})
	r.Get("/logout", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Logout Func"))
	})

	return r
}
