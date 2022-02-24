package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type AdminRoutes struct{}

func (rs AdminRoutes) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world from admin"))
	})

	// Sub routes for indexes actions
	r.Mount("/index", IndexRoutes{}.Routes())
	// Sub routes for document actions
	r.Mount("/document", DocumentRoutes{}.Routes())
	// Configuration API's
	r.Mount("/configure", ConfigurationRoutes{}.Routes())

	// Auth
	r.Get("/login", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Login Func"))
	})
	r.Get("/logout", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Logout Func"))
	})

	return r
}
