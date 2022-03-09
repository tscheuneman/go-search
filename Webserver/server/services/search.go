package services

import (
	"github.com/meilisearch/meilisearch-go"
	"github.com/tscheuneman/go-search/container"
	"github.com/tscheuneman/go-search/data"
)

func Search(query string, search *data.SearchEndpoint, limit int64, offset int64) (resp *meilisearch.SearchResponse, err error) {
	client := container.GetClient()

	index, err := client.GetIndex(search.Index)

	if err != nil {
		return nil, err
	}

	searchRequest := &meilisearch.SearchRequest{
		AttributesToRetrieve:  search.DisplayFields,
		AttributesToHighlight: search.HighlightFields,
		Limit:                 limit,
		Offset:                offset,
	}

	results, err := index.Search(query, searchRequest)

	if err != nil {
		return nil, err
	}

	return results, nil
}
