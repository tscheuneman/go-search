package services

import (
	"github.com/tscheuneman/go-search/container"
	"github.com/tscheuneman/go-search/data"
	"github.com/tscheuneman/go-search/utils"
	"gorm.io/gorm"
)

type ConfigureSearchRequest struct {
	Id                *string   `json:"id,omitempty"`
	Slug              string    `json:"slug,omitempty"`
	DisplayFields     *[]string `json:"display_fields,omitempty"`
	HighlightFields   *[]string `json:"highlight_fields,omitempty"`
	AllowedSortFields *[]string `json:"sort_fields,omitempty"`
	AllowedFacets     *[]string `json:"allowed_facets,omitempty"`
}

func CreateSearchEndpoint(index_slug string, search ConfigureSearchRequest) (res *utils.Status, err error) {
	dbConn := container.GetDb()
	searchEndpoint := data.SearchEndpoint{
		Slug:            search.Slug,
		Index:           index_slug,
		DisplayFields:   *search.DisplayFields,
		HighlightFields: *search.HighlightFields,
		AllowedFacets:   *search.AllowedFacets,
	}

	var result *gorm.DB

	if search.Id == nil {
		result = dbConn.Model(&data.SearchEndpoint{}).Create(&searchEndpoint)
	} else {
		result = dbConn.Model(&data.SearchEndpoint{}).Where("id = ?", search.Id).Updates(&searchEndpoint)
	}

	if result.Error != nil {
		return nil, result.Error
	}

	response := &utils.Status{
		Status:  200,
		Message: "Created Search Index",
	}

	return response, nil
}

func GetSearches(index_slug string) (res interface{}) {
	dbConn := container.GetDb()

	var results []data.SearchEndpoint

	dbConn.Model(&data.SearchEndpoint{}).Where("index = ?", index_slug).Find(&results)

	return results
}

func GetSearch(index_slug string, search_slug string) (res interface{}) {
	dbConn := container.GetDb()

	var result data.SearchEndpoint

	dbConn.Model(&data.SearchEndpoint{}).Where("index = ? AND slug = ?", index_slug, search_slug).Find(&result)

	return result
}
