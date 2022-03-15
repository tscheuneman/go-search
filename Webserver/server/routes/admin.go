package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/tscheuneman/go-search/middleware"
)

type AdminRoutes struct{}

func (rs AdminRoutes) Routes() chi.Router {
	r := chi.NewRouter()

	r.Use(middleware.JwtMiddleware)
	// TODO: Soooome kind of auth middleware.
	// It'll probably be something super basic, probably no roles/ granular permissions

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
