package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/meilisearch/meilisearch-go"
	"github.com/tscheuneman/go-search/container"
	"github.com/tscheuneman/go-search/routes"
)

func main() {
	client := meilisearch.NewClient(meilisearch.ClientConfig{
		Host: "http://meilisearch:7700",
	})

	container.SetClient(client)
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Mount("/", routes.MainRouter{}.Routes())

	http.ListenAndServe(":80", r)
}
