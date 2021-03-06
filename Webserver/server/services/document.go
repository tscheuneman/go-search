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

func GetDocument(index_slug string, document_id string) (resp interface{}, err error) {
	client := container.GetClient()
	index, err := client.GetIndex(index_slug)

	if err != nil {
		return nil, err
	}

	var document interface{}
	update_err := index.GetDocument(document_id, &document)

	if update_err != nil {
		return nil, update_err
	}

	return document, nil
}

func PublishDocuments(index_slug string, request []map[string]interface{}) (resp *meilisearch.Task, err error) {
	client := container.GetClient()
	index, err := client.GetIndex(index_slug)

	if err != nil {
		return nil, err
	}

	task, update_err := index.AddDocuments(request)

	if update_err != nil {
		return nil, update_err
	}

	return task, nil
}

func DeleteDocument(index_slug string, document_id string) (resp *meilisearch.Task, err error) {
	client := container.GetClient()
	index, err := client.GetIndex(index_slug)

	if err != nil {
		return nil, err
	}

	task, update_err := index.DeleteDocument(document_id)

	if update_err != nil {
		return nil, update_err
	}

	return task, nil
}
