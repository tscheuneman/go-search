package controllers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"

	"github.com/tscheuneman/go-search/services"
	"github.com/tscheuneman/go-search/utils"
)

type CreateTokenStruct struct {
	Name string `json:"name,omitempty"`
}

func (a *CreateTokenStruct) Bind(r *http.Request) error {
	return nil
}

func CreateToken(w http.ResponseWriter, r *http.Request) {
	data := &CreateTokenStruct{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, utils.ErrInvalidRequest(err))
		return
	}

	token, err := services.CreateToken(data.Name)
	if err != nil {
		render.Render(w, r, utils.ErrInvalidRequest(err))
		return
	}
	render.JSON(w, r, token)
}

func GetTokens(w http.ResponseWriter, r *http.Request) {

	tokens, err := services.GetAllTokens()

	if err != nil {
		render.Render(w, r, utils.ErrInvalidRequest(err))
		return
	}

	render.JSON(w, r, tokens)
}

func DeleteToken(w http.ResponseWriter, r *http.Request) {
	token_id := chi.URLParam(r, "token_id")

	token, err := services.DeleteToken(token_id)

	if err != nil {
		render.Render(w, r, utils.ErrInvalidRequest(err))
		return
	}

	utils.DeleteCookies(w)

	render.JSON(w, r, token)
}
