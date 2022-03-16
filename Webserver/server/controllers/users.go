package controllers

import (
	"net/http"

	"github.com/go-chi/render"

	"github.com/tscheuneman/go-search/services"
	"github.com/tscheuneman/go-search/utils"
)

func AllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := services.GetAllUsers()

	if err != nil {
		render.Render(w, r, utils.ErrInvalidRequest(err))
		return
	}

	render.JSON(w, r, users)
}
