package controllers

import (
	"errors"
	"net/http"

	"github.com/go-chi/render"

	"github.com/tscheuneman/go-search/container"
	"github.com/tscheuneman/go-search/services"
	"github.com/tscheuneman/go-search/utils"
)

type LoginRequest struct {
	Search services.LoginWebStruct `json:"credentials"`
}

func (a *LoginRequest) Bind(r *http.Request) error {
	return nil
}

func Login(w http.ResponseWriter, r *http.Request) {
	data := &LoginRequest{}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, utils.ErrInvalidRequest(err))
		return
	}

	msg, err := services.LoginWeb(data.Search)

	if err != nil {
		render.Render(w, r, utils.ErrInvalidRequest(errors.New(msg)))
		return
	}

	token, err := utils.GenerateToken(msg)

	if err != nil {
		render.Render(w, r, utils.ErrInvalidRequest(err))
		return
	}

	var useSecure = true

	if container.IS_DEV {
		useSecure = false
	}

	cookie := http.Cookie{
		Name:     container.AUTH_COOKIE,
		Value:    token,
		MaxAge:   259200,
		HttpOnly: true,
		Secure:   useSecure,
		Domain:   container.CLIENT_ORIGIN,
	}

	userCookie := http.Cookie{
		Name:     container.USER_COOKIE,
		Value:    msg,
		MaxAge:   259200,
		HttpOnly: false,
		Secure:   useSecure,
		Domain:   container.CLIENT_ORIGIN,
	}

	http.SetCookie(w, &cookie)
	http.SetCookie(w, &userCookie)

	render.JSON(w, r, utils.Status{
		Status:  200,
		Message: msg,
	})
}
