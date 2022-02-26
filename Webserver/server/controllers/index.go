package controllers

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"

	"github.com/tscheuneman/go-search/services"
	"github.com/tscheuneman/go-search/utils"
)

type IndexRoutes struct{}

type IndexCrud struct {
	Slug string `json:"slug"`
}

type IndexCrudRequest struct {
	*IndexCrud

	ProtectedID string `json:"id"` // override 'id' json to have more control
}

func (a *IndexCrudRequest) Bind(r *http.Request) error {
	if a.IndexCrud == nil {
		return errors.New("missing required Index Crud Fields.")
	}

	return nil
}

func CreateIndex(w http.ResponseWriter, r *http.Request) {
	data := &IndexCrudRequest{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, utils.ErrInvalidRequest(err))
		return
	}

	create, err := services.CreateIndex(data.Slug)

	if err != nil {
		render.Render(w, r, utils.ErrInvalidRequest(err))
		return
	}

	render.Status(r, http.StatusCreated)
	render.Render(w, r, utils.NewTaskResponse(create))
}

func GetIndexes(w http.ResponseWriter, r *http.Request) {

	indexes, err := services.GetAllIndexes()

	if err != nil {
		render.Render(w, r, utils.ErrInvalidRequest(err))
		return
	}

	render.RenderList(w, r, utils.IndexListResponse(indexes))
}

func GetIndex(w http.ResponseWriter, r *http.Request) {
	index_id := chi.URLParam(r, "index_slug")

	index, err := services.GetIndex(index_id)

	if err != nil {
		render.Render(w, r, utils.ErrInvalidRequest(err))
		return
	}

	render.Render(w, r, utils.NewIndexResponse(index))
}

func DeleteIndex(w http.ResponseWriter, r *http.Request) {
	data := &IndexCrudRequest{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, utils.ErrInvalidRequest(err))
		return
	}

	delete, err := services.DeleteIndex(data.Slug)

	if err != nil {
		render.Render(w, r, utils.ErrInvalidRequest(err))
		return
	}

	render.Status(r, http.StatusCreated)
	render.Render(w, r, utils.NewTaskResponse(delete))
}
