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
		authHeader := r.Header.Get(container.AUTH_HEADER)
		if authHeader == "" {
			render.Render(w, r, utils.ErrForbiddenRequest(errors.New("Auth Token Missing")))
			return

		}
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
