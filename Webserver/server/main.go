package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/meilisearch/meilisearch-go"
	"github.com/tscheuneman/go-search/container"
	"github.com/tscheuneman/go-search/data"
	"github.com/tscheuneman/go-search/routes"
	"github.com/tscheuneman/go-search/utils"
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

	fmt.Println("Running preprocess tasks")
	utils.AdminUserPreprocess()

	fmt.Println("Initializing HTTP Server")
	initHttp()
}

func initHttp() {
	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Mount("/", routes.MainRouter{}.Routes())

	http.ListenAndServe(":80", r)
}
