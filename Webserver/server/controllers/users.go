package controllers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"

	"github.com/tscheuneman/go-search/services"
	"github.com/tscheuneman/go-search/utils"
)

type CreateUserStruct struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

func (a *CreateUserStruct) Bind(r *http.Request) error {
	return nil
}

type ChangePwStruct struct {
	Password string `json:"password,omitempty"`
}

func (a *ChangePwStruct) Bind(r *http.Request) error {
	return nil
}

func AllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := services.GetAllUsers()

	if err != nil {
		render.Render(w, r, utils.ErrInvalidRequest(err))
		return
	}

	render.JSON(w, r, users)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	data := &CreateUserStruct{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, utils.ErrInvalidRequest(err))
		return
	}

	user, err := services.CreateUser(data.Username, data.Password)

	if err != nil {
		render.Render(w, r, utils.ErrInvalidRequest(err))
		return
	}

	render.JSON(w, r, user)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	user_id := chi.URLParam(r, "user_id")

	user, err := services.DeleteUser(user_id)

	if err != nil {
		render.Render(w, r, utils.ErrInvalidRequest(err))
		return
	}

	render.JSON(w, r, user)
}

func ChangePassword(w http.ResponseWriter, r *http.Request) {
	data := &ChangePwStruct{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, utils.ErrInvalidRequest(err))
		return
	}

	user_id := chi.URLParam(r, "user_id")

	user, err := services.ChangePassword(user_id, data.Password)

	if err != nil {
		render.Render(w, r, utils.ErrInvalidRequest(err))
		return
	}

	render.JSON(w, r, user)
}
