package services

import (
	"github.com/meilisearch/meilisearch-go"
	"github.com/tscheuneman/go-search/container"
	"github.com/tscheuneman/go-search/data"
)

func Search(query string, search *data.SearchEndpoint, limit int64, offset int64, FilterConstructor []string) (resp *meilisearch.SearchResponse, err error) {
	client := container.GetClient()

	index, err := client.GetIndex(search.Index)

	if err != nil {
		return nil, err
	}

	var AllowedFacets []string

	if search.AllowedFacets != nil && len(search.AllowedFacets) > 0 && search.AllowedFacets[0] != "" {
		AllowedFacets = search.AllowedFacets
	} else {
		AllowedFacets = nil
	}

	var Filters []string

	if len(FilterConstructor) > 0 {
		Filters = FilterConstructor
	} else {
		Filters = nil
	}

	searchRequest := &meilisearch.SearchRequest{
		AttributesToRetrieve:  search.DisplayFields,
		AttributesToHighlight: search.HighlightFields,
		FacetsDistribution:    AllowedFacets,
		Limit:                 limit,
		Offset:                offset,
		Filter:                Filters,
	}

	results, err := index.Search(query, searchRequest)

	if err != nil {
		return nil, err
	}

	return results, nil
}
