package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/meilisearch/meilisearch-go"
	"github.com/tscheuneman/go-search/container"
	"github.com/tscheuneman/go-search/data"
	"github.com/tscheuneman/go-search/routes"
)

func main() {
	meliUrl := os.Getenv(container.MEILI_URL)
	dbHost := os.Getenv(container.DB_HOST)
	dbUser := os.Getenv(container.DB_USER)
	dbPw := os.Getenv(container.DB_PASSWORD)

	if meliUrl == "" || dbHost == "" || dbUser == "" || dbPw == "" {
		panic("Required Env Variables don't exist")
	}

	client := meilisearch.NewClient(meilisearch.ClientConfig{
		Host: meliUrl,
	})

	dsn := "host=" + dbHost + " user=" + dbUser + " password=" + dbPw

	dbConn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Couldn't initalize DB")
		panic(err)
	}

	container.SetDb(dbConn)
	container.SetClient(client)

	// DB Migrations should probably live somewhere else.  This is fine for now though
	dbConn.AutoMigrate(&data.User{}, &data.SearchEndpoint{})

	fmt.Println("Initializing HTTP Server")
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
