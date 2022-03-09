package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/tscheuneman/go-search/controllers"
)

type SearchRoutes struct{}

func (rs SearchRoutes) Routes() chi.Router {
	r := chi.NewRouter()

	// TODO: Soooome kind of auth middleware.
	// It'll probably be something super basic, probably no roles/ granular permissions

	r.Get("/{search_slug}", controllers.Search)

	return r
}
