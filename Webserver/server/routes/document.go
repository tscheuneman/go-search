package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/tscheuneman/go-search/controllers"
)

type DocumentRoutes struct{}

func (rs DocumentRoutes) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", controllers.GetDocuments)
	r.Post("/", controllers.CreateUpdateDocuments)
	r.Delete("/{document_id}", controllers.DeleteDocument)

	r.Get("/{document_id}", controllers.GetDocument)

	return r
}
