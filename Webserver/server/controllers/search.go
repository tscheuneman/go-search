package controllers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"

	"github.com/tscheuneman/go-search/services"
	"github.com/tscheuneman/go-search/utils"
)

func Search(w http.ResponseWriter, r *http.Request) {
	search_slug := chi.URLParam(r, "search_slug")
	query := r.URL.Query().Get("q")

	limit := utils.QueryParamToInt64(r, "limit", 100)
	offset := utils.QueryParamToInt64(r, "offset", 0)

	search, err := services.GetSearch(search_slug)

	if err != nil {
		render.Render(w, r, utils.ErrInvalidRequest(err))
		return
	}

	search_results, err := services.Search(query, search, limit, offset)

	if err != nil {
		render.Render(w, r, utils.ErrInvalidRequest(err))
		return
	}

	render.JSON(w, r, search_results)
}
