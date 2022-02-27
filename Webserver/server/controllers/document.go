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
	Documents []services.CreateCardRequest `json:"documents,omitempty"`
}

func (a *DocumentCrudRequest) Bind(r *http.Request) error {
	for _, val := range a.Documents {
		if val.Id == nil {
			return errors.New("All documents must have an ID")
		}
	}
	if a.Documents == nil {
		return errors.New("missing required to create documents")
	}
	return nil
}

func GetDocuments(w http.ResponseWriter, r *http.Request) {
	index_slug := chi.URLParam(r, "index_slug")

	limit := utils.QueryParamToInt64(r, "limit", 100)
	offset := utils.QueryParamToInt64(r, "offset", 0)

	documents, err := services.GetAllDocuments(index_slug, limit, offset)

	if err != nil {
		render.Render(w, r, utils.ErrInvalidRequest(err))
		return
	}

	render.JSON(w, r, documents)
}

func CreateUpdateDocuments(w http.ResponseWriter, r *http.Request) {
	data := &DocumentCrudRequest{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, utils.ErrInvalidRequest(err))
		return
	}

	index_slug := chi.URLParam(r, "index_slug")

	documents, err := services.PublishDocuments(index_slug, data.Documents)

	if err != nil {
		render.Render(w, r, utils.ErrInvalidRequest(err))
		return
	}

	render.Status(r, http.StatusCreated)
	render.Render(w, r, utils.NewTaskResponse(documents))
}
