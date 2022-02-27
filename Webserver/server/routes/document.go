package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/tscheuneman/go-search/controllers"
)

type DocumentRoutes struct{}

func (rs DocumentRoutes) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", controllers.GetDocuments)

	return r
}
