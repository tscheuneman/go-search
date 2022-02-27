package controllers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"

	"github.com/tscheuneman/go-search/services"
	"github.com/tscheuneman/go-search/utils"
)

type ConfigureGlobalsCrudRequest struct {
	Config services.ConfigureGlobals `json:"config,omitempty"`
}

func (a *ConfigureGlobalsCrudRequest) Bind(r *http.Request) error {
	return nil
}

func ConfigureGlobals(w http.ResponseWriter, r *http.Request) {
	data := &ConfigureGlobalsCrudRequest{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, utils.ErrInvalidRequest(err))
		return
	}

	index_slug := chi.URLParam(r, "index_slug")

	task, err := services.SetGlobalConfig(index_slug, data.Config)

	if err != nil {
		render.Render(w, r, utils.ErrInvalidRequest(err))
		return
	}

	render.JSON(w, r, task)
}

func GetGlobals(w http.ResponseWriter, r *http.Request) {
	index_slug := chi.URLParam(r, "index_slug")

	settings, err := services.GetGlobalConfig(index_slug)

	if err != nil {
		render.Render(w, r, utils.ErrInvalidRequest(err))
		return
	}

	render.JSON(w, r, settings)
}
