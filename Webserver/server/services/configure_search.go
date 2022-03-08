package services

import (
	"github.com/tscheuneman/go-search/container"
	"github.com/tscheuneman/go-search/data"
)

type ConfigureSearchRequest struct {
	Id                *string   `json:"id,omitempty"`
	Slug              string    `json:"slug,omitempty"`
	DisplayFields     *[]string `json:"display_fields,omitempty"`
	HighlightFields   *[]string `json:"highlight_fields,omitempty"`
	AllowedSortFields *[]string `json:"sort_fields,omitempty"`
	AllowedFacets     *[]string `json:"allowed_facets,omitempty"`
}

func CreateSearchEndpoint(index_slug string, search ConfigureSearchRequest) (res bool) {
	dbConn := container.GetDb()
	searchEndpoint := data.SearchEndpoint{
		Slug:            search.Slug,
		Index:           index_slug,
		DisplayFields:   *search.DisplayFields,
		HighlightFields: *search.HighlightFields,
		AllowedFacets:   *search.AllowedFacets,
	}
	if search.Id != nil {
		dbConn.Create(&searchEndpoint)
	} else {
		dbConn.Model(&data.SearchEndpoint{}).Where("id = ?", search.Id).Updates(&searchEndpoint)
	}

	return true
}

func GetSearches(index_slug string) (res interface{}) {
	dbConn := container.GetDb()

	searches := dbConn.Where("index = ?", index_slug).Find(&data.SearchEndpoint{})

	return searches
}

func GetSearch(index_slug string, search_slug string) (res interface{}) {
	dbConn := container.GetDb()

	search := dbConn.Where("index = ? AND slug = ?", index_slug, search_slug).Find(&data.SearchEndpoint{})

	return search
}
