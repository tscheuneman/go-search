package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/tscheuneman/go-search/controllers"
)

type MainRouter struct{}

func (rs MainRouter) Routes() chi.Router {
	r := chi.NewRouter()

	r.Mount("/admin", AdminRoutes{}.Routes())
	r.Mount("/search", SearchRoutes{}.Routes())

	r.Post("/login", controllers.Login)

	return r
}
