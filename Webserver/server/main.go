package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/meilisearch/meilisearch-go"
	"github.com/tscheuneman/go-search/container"
	"github.com/tscheuneman/go-search/routes"
)

func main() {
	meliUrl := os.Getenv(container.MEILI_URL)
	if meliUrl == "" {
		log.Println("Required Env Variables don't exist")
		os.Exit(1)
	}

	client := meilisearch.NewClient(meilisearch.ClientConfig{
		Host: meliUrl,
	})

	container.SetClient(client)

	initHttp()
}

func initHttp() {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Mount("/", routes.MainRouter{}.Routes())

	http.ListenAndServe(":80", r)
}
