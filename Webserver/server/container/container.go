package container

import (
	"os"
	"sync"

	lru "github.com/hashicorp/golang-lru"
	"github.com/meilisearch/meilisearch-go"

	"github.com/tscheuneman/go-search/implementor"
	"gorm.io/gorm"
)

const DEFAULT_USER = "DEFAULT_USER"
const MEILI_URL = "MEILI_URL"
const SERVICE_PORT = "SERVICE_PORT"

// DB
const DB_HOST = "DB_HOST"
const DB_USER = "DB_USER"
const DB_PASSWORD = "DB_PASSWORD"
const CLIENT_ORIGIN_ENV = "CLIENT_ORIGIN_ENV"
const JWT_KEY_ENV = "JWT_KEY"

// Setable Values
var JWT_KEY string
var CLIENT_ORIGIN string

var IS_DEV bool = os.Getenv("ENV") == "DEV"

var dbSingleton sync.Once
var clientSingleton sync.Once
var cacheClientSingleton sync.Once

var clientConnection *meilisearch.Client
var dbConnection *gorm.DB
var cacheClient *implementor.CacheClient

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

func SetCacheClient() {
	cacheClientSingleton.Do(func() {
		client, err := lru.New(50)

		if err != nil {
			panic("Couldn't initalize cache")
		}
		cacheClient = &implementor.CacheClient{
			Client: client,
			Ttl:    3600,
		}
	})
}

func GetCacheClient() *implementor.CacheClient {
	return cacheClient
}
