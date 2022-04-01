package middleware

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"

	"github.com/tscheuneman/go-search/utils"
)

func JwtMiddleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)

		/*
			authCookie, err := r.Cookie(container.AUTH_COOKIE)
			if err != nil {
				render.Render(w, r, utils.ErrForbiddenRequest(errors.New("Auth Token Missing")))
				return
			}
			validToken, user_id := utils.ValidateToken(authCookie.Value)
			if !validToken {
				render.Render(w, r, utils.ErrForbiddenRequest(errors.New("Invalid Auth Token")))
				return
			}

			ctx := r.Context()
			ctx = context.WithValue(ctx, "UserID", user_id)
			r = r.WithContext(ctx)

			next.ServeHTTP(w, r)
		*/
	}
	return http.HandlerFunc(fn)
}

func UserOnlyAuthMiddleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		userId := r.Context().Value("UserID")
		fmt.Println(userId)

		paramUserId := chi.URLParam(r, "user_id")
		fmt.Println(paramUserId)
		if paramUserId != "" && userId != paramUserId {
			render.Render(w, r, utils.ErrNotAuthorizedRequest(errors.New("UserId Mismatch")))
			return
		}

		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
