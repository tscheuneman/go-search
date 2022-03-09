package routes

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-chi/chi/v5"
)

type MainRouter struct{}

func (rs MainRouter) Routes() chi.Router {
	r := chi.NewRouter()

	workDir, _ := os.Getwd()
	filesDir := http.Dir(filepath.Join(workDir, "../", "client"))
	fmt.Println(filesDir)

	FileServer(r, "/", filesDir)

	r.Mount("/admin", AdminRoutes{}.Routes())
	r.Mount("/search", SearchRoutes{}.Routes())

	return r
}

func FileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit any URL parameters.")
	}

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, func(w http.ResponseWriter, r *http.Request) {
		rctx := chi.RouteContext(r.Context())
		pathPrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")
		fs := http.StripPrefix(pathPrefix, http.FileServer(root))
		fs.ServeHTTP(w, r)
	})
}
