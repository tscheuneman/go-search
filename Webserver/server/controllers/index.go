package controllers

import (
	"errors"
	"net/http"

	"github.com/go-chi/render"
	"github.com/meilisearch/meilisearch-go"

	"github.com/tscheuneman/go-search/container"
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

	client := container.GetClient()

	index, _ := client.GetIndex(data.Slug)

	if index != nil {
		render.Render(w, r, utils.HttpError("Index exists", 409))
		return
	}

	config := &meilisearch.IndexConfig{
		Uid:        data.Slug,
		PrimaryKey: "id",
	}

	create, err := client.CreateIndex(config)

	if err != nil {
		render.Render(w, r, utils.ErrInvalidRequest(err))
		return
	}

	render.Status(r, http.StatusCreated)
	render.Render(w, r, utils.NewTaskResponse(create))
}

func GetIndexes(w http.ResponseWriter, r *http.Request) {
	client := container.GetClient()
	indexes, err := client.GetAllIndexes()

	if err != nil {
		render.Render(w, r, utils.ErrInvalidRequest(err))
		return
	}

	render.RenderList(w, r, utils.IndexListResponse(indexes))
}

func DeleteIndex(w http.ResponseWriter, r *http.Request) {
	data := &IndexCrudRequest{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, utils.ErrInvalidRequest(err))
		return
	}

	client := container.GetClient()

	delete, err := client.DeleteIndex(data.Slug)

	if err != nil {
		render.Render(w, r, utils.ErrInvalidRequest(err))
		return
	}
	render.Status(r, http.StatusCreated)
	render.Render(w, r, utils.NewTaskResponse(delete))
}
