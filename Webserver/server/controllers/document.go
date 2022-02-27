package controllers

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"

	"github.com/tscheuneman/go-search/services"
	"github.com/tscheuneman/go-search/utils"
)

type Documentroutes struct{}

type DocumentCrud struct {
	Slug string `json:"slug"`
}

type DocumentCrudRequest struct {
	*DocumentCrud

	ProtectedID string `json:"id"` // override 'id' json to have more control
}

func (a *DocumentCrudRequest) Bind(r *http.Request) error {
	if a.DocumentCrud == nil {
		return errors.New("missing required Document Crud Fields.")
	}

	return nil
}

func GetDocuments(w http.ResponseWriter, r *http.Request) {
	index_slug := chi.URLParam(r, "index_slug")

	limit := utils.QueryParamToInt64(r, "limit", 100)
	offset := utils.QueryParamToInt64(r, "offset", 100)

	documents, err := services.GetAllDocuments(index_slug, limit, offset)

	if err != nil {
		render.Render(w, r, utils.ErrInvalidRequest(err))
		return
	}

	render.JSON(w, r, documents)

}
