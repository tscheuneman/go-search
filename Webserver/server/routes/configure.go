package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/tscheuneman/go-search/controllers"
)

type ConfigurationRoutes struct{}

func (rs ConfigurationRoutes) Routes() chi.Router {
	r := chi.NewRouter()

	r.Post("/globals", controllers.ConfigureGlobals)
	r.Get("/globals", controllers.GetGlobals)

	r.Post("/search", controllers.ConfigureSearch)
	r.Get("/search", controllers.GetSearches)
	r.Get("/search/{search_slug}", controllers.GetSearch)
	r.Delete("/search/{search_slug}", controllers.DeleteSearch)
	return r
}
