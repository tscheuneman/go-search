package container

import (
	"sync"

	"github.com/meilisearch/meilisearch-go"
	"gorm.io/gorm"
)

const DEFAULT_USER = "DEFAULT_USER"

const MEILI_URL = "MEILI_URL"

// DB
const DB_HOST = "DB_HOST"
const DB_USER = "DB_USER"
const DB_PASSWORD = "DB_PASSWORD"

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
