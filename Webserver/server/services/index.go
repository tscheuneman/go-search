package services

import (
	"errors"

	"github.com/meilisearch/meilisearch-go"

	"github.com/tscheuneman/go-search/container"
)

func CreateIndex(index_slug string) (resp *meilisearch.Task, err error) {
	client := container.GetClient()
	index, _ := client.GetIndex(index_slug)

	if index != nil {
		return nil, errors.New("Index already exists")
	}

	config := &meilisearch.IndexConfig{
		Uid:        index_slug,
		PrimaryKey: "id",
	}

	create, err := client.CreateIndex(config)

	if err != nil {
		return nil, err
	}

	return create, nil
}

func GetAllIndexes() (resp []*meilisearch.Index, err error) {
	client := container.GetClient()
	indexes, err := client.GetAllIndexes()

	if err != nil {
		return nil, err
	}
	return indexes, nil
}

func DeleteIndex(index_slug string) (resp *meilisearch.Task, err error) {
	client := container.GetClient()

	delete, err := client.DeleteIndex(index_slug)

	if err != nil {
		return nil, err
	}

	return delete, nil
}
