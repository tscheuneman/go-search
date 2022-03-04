package utils

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/render"
)

type ErrResponse struct {
	Err            error `json:"-"` // low-level runtime error
	HTTPStatusCode int   `json:"-"` // http response status code

	StatusText string `json:"status"`          // user-level status message
	AppCode    int64  `json:"code,omitempty"`  // application-specific error code
	ErrorText  string `json:"error,omitempty"` // application-level error message, for debugging
}

func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

func ErrInvalidRequest(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 400,
		StatusText:     "Invalid request.",
		ErrorText:      err.Error(),
	}
}

func HttpError(errorText string, status int) render.Renderer {
	err := errors.New(errorText)
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: status,
		StatusText:     errorText,
		ErrorText:      err.Error(),
	}
}

func QueryParamToInt64(req *http.Request, name string, default_val int64) int64 {
	param := req.URL.Query().Get(name)
	result, err := strconv.ParseInt(param, 10, 64)
	if err != nil {
		return default_val
	}
	return result
}

func SetResponseHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Access-Control-Allow-Origin", "*")

		next.ServeHTTP(rw, r)
	})
}
