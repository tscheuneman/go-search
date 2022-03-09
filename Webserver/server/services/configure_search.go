package services

import (
	"fmt"

	"github.com/lib/pq"
	"github.com/tscheuneman/go-search/container"
	"github.com/tscheuneman/go-search/data"
	"github.com/tscheuneman/go-search/utils"
	"gorm.io/gorm"
)

type ConfigureSearchRequest struct {
	Id              *string  `json:"id,omitempty"`
	Slug            string   `json:"slug,omitempty"`
	DisplayFields   []string `json:"display_fields,omitempty"`
	HighlightFields []string `json:"highlight_fields,omitempty"`
	AllowedFacets   []string `json:"allowed_facets,omitempty"`
}

func CreateSearchEndpoint(index_slug string, search ConfigureSearchRequest) (res *utils.Status, err error) {
	dbConn := container.GetDb()

	var result *gorm.DB

	if search.Id == nil {
		result = dbConn.Model(&data.SearchEndpoint{}).Create(&data.SearchEndpoint{
			DisplayFields:   pq.StringArray(search.DisplayFields),
			HighlightFields: pq.StringArray(search.HighlightFields),
			AllowedFacets:   pq.StringArray(search.AllowedFacets),
			Index:           index_slug,
			Slug:            search.Slug,
		})
	} else {
		result = dbConn.Model(&data.SearchEndpoint{}).Where("id = ?", search.Id).Updates(&data.SearchEndpoint{
			DisplayFields:   pq.StringArray(search.DisplayFields),
			HighlightFields: pq.StringArray(search.HighlightFields),
			AllowedFacets:   pq.StringArray(search.AllowedFacets),
		})
	}

	if result.Error != nil {
		fmt.Println(result.Error)
		return nil, result.Error
	}

	response := &utils.Status{
		Status:  200,
		Message: "Created Search Index",
	}

	return response, nil
}

func GetSearches(index_slug string) (res *[]data.SearchEndpoint, err error) {
	dbConn := container.GetDb()

	var results *[]data.SearchEndpoint

	dbResult := dbConn.Model(&data.SearchEndpoint{}).Where("index = ?", index_slug).Find(&results)

	if dbResult.Error != nil {
		return nil, dbResult.Error
	}

	return results, nil
}

func GetSearch(search_slug string) (res *data.SearchEndpoint, err error) {
	dbConn := container.GetDb()

	var result *data.SearchEndpoint

	dbResult := dbConn.Model(&data.SearchEndpoint{}).Where("slug = ?", search_slug).Find(&result)

	if dbResult.Error != nil {
		return nil, dbResult.Error
	}

	return result, nil
}
