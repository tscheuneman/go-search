package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/tscheuneman/go-search/controllers"
)

type UserRoutes struct{}

func (rs UserRoutes) Routes() chi.Router {
	r := chi.NewRouter()

	// TODO: Soooome kind of auth middleware.
	// It'll probably be something super basic, probably no roles/ granular permissions

	r.Get("/", controllers.AllUsers)
	r.Post("/", controllers.CreateUser)
	r.Delete("/{user_id}", controllers.DeleteUser)

	return r
}
