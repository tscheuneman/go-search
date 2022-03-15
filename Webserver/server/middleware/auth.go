package middleware

import (
	"errors"
	"net/http"

	"github.com/go-chi/render"

	"github.com/tscheuneman/go-search/container"
	"github.com/tscheuneman/go-search/utils"
)

func JwtMiddleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		authCookie, err := r.Cookie(container.AUTH_COOKIE)
		if err != nil {
			render.Render(w, r, utils.ErrForbiddenRequest(errors.New("Auth Token Missing")))
			return

		}
		validToken := utils.ValidateToken(authCookie.Value)
		if !validToken {
			render.Render(w, r, utils.ErrForbiddenRequest(errors.New("Invalid Auth Token")))
			return
		}

		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
