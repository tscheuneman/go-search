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
	Id                *string  `json:"id,omitempty"`
	Slug              string   `json:"slug,omitempty"`
	DisplayFields     []string `json:"display_fields,omitempty"`
	HighlightFields   []string `json:"highlight_fields,omitempty"`
	CombinationFacets []string `json:"combination_facets,omitempty"`
	AllowedFacets     []string `json:"allowed_facets,omitempty"`
}

func CreateSearchEndpoint(index_slug string, search ConfigureSearchRequest) (res *utils.Status, err error) {
	dbConn := container.GetDb()
	cacheClient := container.GetCacheClient()

	var result *gorm.DB

	if search.Id == nil {
		result = dbConn.Model(&data.SearchEndpoint{}).Create(&data.SearchEndpoint{
			DisplayFields:     pq.StringArray(search.DisplayFields),
			HighlightFields:   pq.StringArray(search.HighlightFields),
			AllowedFacets:     pq.StringArray(search.AllowedFacets),
			CombinationFacets: pq.StringArray(search.CombinationFacets),
			Index:             index_slug,
			Slug:              search.Slug,
		})
	} else {
		result = dbConn.Model(&data.SearchEndpoint{}).Where("id = ?", search.Id).Updates(&data.SearchEndpoint{
			DisplayFields:     pq.StringArray(search.DisplayFields),
			HighlightFields:   pq.StringArray(search.HighlightFields),
			AllowedFacets:     pq.StringArray(search.AllowedFacets),
			CombinationFacets: pq.StringArray(search.CombinationFacets),
		})
	}

	if result.Error != nil {
		fmt.Println(result.Error)
		return nil, result.Error
	}

	cacheClient.RemoveCacheItem("search_" + search.Slug)

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

func GetSearch(search_slug string) (res interface{}, err error) {
	dbConn := container.GetDb()
	cacheClient := container.GetCacheClient()

	result, err := cacheClient.Resolve("search_"+search_slug, func() (interface{}, error) {
		var result *data.SearchEndpoint

		dbResult := dbConn.Model(&data.SearchEndpoint{}).Where("slug = ?", search_slug).Find(&result)

		if dbResult.Error != nil {
			return nil, dbResult.Error
		}

		return result, nil
	})

	if err != nil {
		return nil, err
	}

	return result, nil
}

func DeleteSearch(index_slug string, search_slug string) (res *utils.Status, err error) {
	dbConn := container.GetDb()

	dbResult := dbConn.Where("index = ? AND slug = ?", index_slug, search_slug).Delete(&data.SearchEndpoint{})

	if dbResult.Error != nil {
		return nil, dbResult.Error
	}

	statusMessage := &utils.Status{
		Status:  200,
		Message: "Search Deleted",
	}

	return statusMessage, nil
}
