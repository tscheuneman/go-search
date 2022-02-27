package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/tscheuneman/go-search/controllers"
)

type IndexRoutes struct{}

func (rs IndexRoutes) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", controllers.GetIndexes)
	r.Post("/", controllers.CreateIndex)
	r.Delete("/", controllers.DeleteIndex)

	r.Get("/{index_slug}", controllers.GetIndex)
	// Sub routes for document actions
	r.Mount("/{index_slug}/document", DocumentRoutes{}.Routes())

	return r
}
