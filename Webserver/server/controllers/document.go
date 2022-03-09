package controllers

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"

	"github.com/tscheuneman/go-search/services"
	"github.com/tscheuneman/go-search/utils"
)

type DocumentCrud struct {
	Slug string `json:"slug"`
}

type DocumentCrudRequest struct {
	Documents []map[string]interface{} `json:"documents,omitempty"`
}

func (a *DocumentCrudRequest) Bind(r *http.Request) error {
	for _, val := range a.Documents {
		if val["id"] == nil {
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

	limit := utils.QueryParamToInt64(r, "limit", 20)
	offset := utils.QueryParamToInt64(r, "offset", 0)

	documents, err := services.GetAllDocuments(index_slug, limit, offset)

	if err != nil {
		render.Render(w, r, utils.ErrInvalidRequest(err))
		return
	}

	render.JSON(w, r, documents)
}

func GetDocument(w http.ResponseWriter, r *http.Request) {
	index_slug := chi.URLParam(r, "index_slug")
	document_id := chi.URLParam(r, "document_id")

	document, err := services.GetDocument(index_slug, document_id)

	if err != nil {
		render.Render(w, r, utils.ErrInvalidRequest(err))
		return
	}

	render.JSON(w, r, document)
}

func CreateUpdateDocuments(w http.ResponseWriter, r *http.Request) {
	data := &DocumentCrudRequest{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, utils.ErrInvalidRequest(err))
		return
	}

	index_slug := chi.URLParam(r, "index_slug")

	publish_task, err := services.PublishDocuments(index_slug, data.Documents)

	if err != nil {
		render.Render(w, r, utils.ErrInvalidRequest(err))
		return
	}

	render.Status(r, http.StatusCreated)
	render.Render(w, r, utils.NewTaskResponse(publish_task))
}

func DeleteDocument(w http.ResponseWriter, r *http.Request) {
	index_slug := chi.URLParam(r, "index_slug")
	document_id := chi.URLParam(r, "document_id")

	delete_task, err := services.DeleteDocument(index_slug, document_id)

	if err != nil {
		render.Render(w, r, utils.ErrInvalidRequest(err))
		return
	}

	render.Status(r, http.StatusCreated)
	render.Render(w, r, utils.NewTaskResponse(delete_task))
}
