package services

import (
	"github.com/meilisearch/meilisearch-go"

	"github.com/tscheuneman/go-search/container"
)

func GetAllDocuments(index_slug string, limit int64, offset int64) (resp []interface{}, err error) {
	client := container.GetClient()
	index, err := client.GetIndex(index_slug)

	if err != nil {
		return nil, err
	}

	var documents []interface{}

	document_request := &meilisearch.DocumentsRequest{
		Limit:  limit,
		Offset: offset,
	}

	docErr := index.GetDocuments(document_request, &documents)

	if err != nil {
		return nil, docErr
	}

	return documents, nil
}
