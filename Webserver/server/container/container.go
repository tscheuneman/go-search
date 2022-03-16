package container

import (
	"os"
	"sync"

	"github.com/meilisearch/meilisearch-go"
	"gorm.io/gorm"
)

const DEFAULT_USER = "DEFAULT_USER"
const MEILI_URL = "MEILI_URL"
const SERVICE_PORT = "SERVICE_PORT"

// DB
const DB_HOST = "DB_HOST"
const DB_USER = "DB_USER"
const DB_PASSWORD = "DB_PASSWORD"

// Setable Values
var JWT_KEY string = "JWT_KEY"
var CLIENT_ORIGIN string = "localhost"

var IS_DEV bool = os.Getenv("ENV") == "DEV"

var dbSingleton sync.Once
var clientSingleton sync.Once

var clientConnection *meilisearch.Client
var dbConnection *gorm.DB

func SetDb(dbCon *gorm.DB) {
	dbSingleton.Do(func() {
		dbConnection = dbCon
	})
}

func GetDb() *gorm.DB {
	return dbConnection
}

func SetClient(clientCon *meilisearch.Client) {
	clientSingleton.Do(func() {
		clientConnection = clientCon
	})
}

func GetClient() *meilisearch.Client {
	return clientConnection
}
