package services

import (
	"fmt"

	"github.com/meilisearch/meilisearch-go"
	"github.com/tscheuneman/go-search/container"
)

type ConfigureGlobals struct {
	DisplayedFields  *[]string `json:"displayed_fields,omitempty"`
	SearchableFields *[]string `json:"searchable_fields,omitempty"`
	FilterableFields *[]string `json:"filterable_fields,omitempty"`
	SortableFields   *[]string `json:"sortable_fields,omitempty"`
}

func SetGlobalConfig(index_slug string, data ConfigureGlobals) (resp *meilisearch.Task, err error) {
	client := container.GetClient()
	index, err := client.GetIndex(index_slug)

	if err != nil {
		return nil, err
	}

	setting_config := &meilisearch.Settings{}

	if data.DisplayedFields != nil {
		setting_config.DisplayedAttributes = *data.DisplayedFields
	}
	if data.SearchableFields != nil {
		setting_config.SearchableAttributes = *data.SearchableFields
	}
	if data.FilterableFields != nil {
		setting_config.FilterableAttributes = *data.FilterableFields
	}
	if data.SortableFields != nil {
		setting_config.SortableAttributes = *data.SortableFields
	}

	fmt.Println(setting_config)

	task, update_err := index.UpdateSettings(setting_config)

	if update_err != nil {
		return nil, update_err
	}

	return task, nil
}

func GetGlobalConfig(index_slug string) (resp *meilisearch.Settings, err error) {
	client := container.GetClient()
	index, err := client.GetIndex(index_slug)

	if err != nil {
		return nil, err
	}

	settings, update_err := index.GetSettings()

	if update_err != nil {
		return nil, update_err
	}

	return settings, nil
}
