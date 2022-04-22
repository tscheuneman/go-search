package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/tscheuneman/go-search/controllers"
)

type TokenRoutes struct{}

func (rs TokenRoutes) Routes() chi.Router {
	r := chi.NewRouter()

	r.Post("/", controllers.CreateToken)
	r.Get("/", controllers.GetTokens)
	r.Delete("/{token_id}", controllers.DeleteToken)

	return r
}
