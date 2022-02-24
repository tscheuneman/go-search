package container

import (
	"sync"

	"github.com/meilisearch/meilisearch-go"
)

var dbSingleton sync.Once
var clientSingleton sync.Once

var clientConnection *meilisearch.Client

func SetDb() {

}

func GetDb() {

}

func SetClient(clientCon *meilisearch.Client) {
	clientSingleton.Do(func() {
		clientConnection = clientCon
	})
}

func GetClient() *meilisearch.Client {
	return clientConnection
}
