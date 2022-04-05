package controllers

import (
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"

	"github.com/tscheuneman/go-search/data"
	"github.com/tscheuneman/go-search/services"
	"github.com/tscheuneman/go-search/utils"
)

func Search(w http.ResponseWriter, r *http.Request) {
	search_slug := chi.URLParam(r, "search_slug")
	query := r.URL.Query().Get("q")

	limit := utils.QueryParamToInt64(r, "limit", 20)
	offset := utils.QueryParamToInt64(r, "offset", 0)

	searchInterface, err := services.GetSearch(search_slug)

	search, ok := searchInterface.(*data.SearchEndpoint)

	if !ok {
		render.Render(w, r, utils.ErrInvalidRequest(err))
		return
	}

	var FilterConstructor []string

	for i := 0; i < len(search.AllowedFacets); i++ {
		searchParam := r.URL.Query().Get(search.AllowedFacets[i])
		if searchParam != "" {
			split := strings.Split(searchParam, ",")

			for x := 0; x < len(split); x++ {
				FilterConstructor = append(FilterConstructor, search.AllowedFacets[i]+" = "+split[x])
			}
		}
	}

	if err != nil {
		render.Render(w, r, utils.ErrInvalidRequest(err))
		return
	}

	search_results, err := services.Search(query, search, limit, offset, FilterConstructor)

	if err != nil {
		render.Render(w, r, utils.ErrInvalidRequest(err))
		return
	}

	render.JSON(w, r, search_results)
}
