package middleware

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/go-chi/render"

	"github.com/tscheuneman/go-search/container"
	"github.com/tscheuneman/go-search/utils"
)

func JwtMiddleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		authCookie, err := r.Cookie(container.AUTH_HEADER)
		if err != nil {
			render.Render(w, r, utils.ErrForbiddenRequest(errors.New("Auth Token Missing")))
			return

		}
		fmt.Println(authCookie)

		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
