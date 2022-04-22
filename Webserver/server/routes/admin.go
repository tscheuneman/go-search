package routes

import (
	"github.com/go-chi/chi/v5"

	"github.com/tscheuneman/go-search/middleware"
)

type AdminRoutes struct{}

func (rs AdminRoutes) Routes() chi.Router {
	r := chi.NewRouter()

	r.Use(middleware.JwtMiddleware)

	// Sub routes for indexes actions
	r.Mount("/index", IndexRoutes{}.Routes())
	r.Mount("/users", UserRoutes{}.Routes())

	r.Mount("/token", TokenRoutes{}.Routes())

	return r
}
