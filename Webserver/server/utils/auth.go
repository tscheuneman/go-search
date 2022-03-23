package utils

import (
	"net/http"

	"github.com/tscheuneman/go-search/container"
)

func DeleteCookies(w http.ResponseWriter) {
	var useSecure = true

	if container.IS_DEV {
		useSecure = false
	}

	cookie := http.Cookie{
		Name:     container.AUTH_COOKIE,
		Value:    "",
		MaxAge:   -1,
		HttpOnly: true,
		Secure:   useSecure,
		Path:     "/",
		Domain:   container.CLIENT_ORIGIN,
	}

	userCookie := http.Cookie{
		Name:     container.USER_COOKIE,
		Value:    "",
		MaxAge:   -1,
		HttpOnly: false,
		Secure:   useSecure,
		Path:     "/",
		Domain:   container.CLIENT_ORIGIN,
	}

	http.SetCookie(w, &cookie)
	http.SetCookie(w, &userCookie)
}
