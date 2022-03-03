package services

type ConfigureSearchRequest struct {
	Slug              string    `json:"slug,omitempty"`
	DisplayFields     *[]string `json:"display_fields,omitempty"`
	HighlightFields   *[]string `json:"highlight_fields,omitempty"`
	AllowedSortFields *[]string `json:"sort_fields,omitempty"`
	AllowedFacets     *[]string `json:"allowed_facets,omitempty"`
}

func ConfigureSearch(index_slug string, data ConfigureSearchRequest) (resp interface{}, err error) {

	return nil, nil
}
