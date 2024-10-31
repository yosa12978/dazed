package router

import (
	"fmt"
	"io"
	"net/http"

	"github.com/yosa12978/dazed/internal/middleware"
)

func NewRouter(opts ...optionFunc) http.Handler {
	options := newOptions(opts...)
	options.logger.Debug("initializing router")
	router := http.NewServeMux()
	addV1Routes(router, options)
	handler := middleware.Composition(
		router,
		middleware.Logger(options.logger),
		middleware.StripSlash(),
		middleware.Recovery(options.logger),
	)
	return handler
}

func addV1Routes(router *http.ServeMux, options routerOptions) {
	options.logger.Debug("adding v1 routes")
	router.HandleFunc("GET /{$}", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "new router")
	})
	router.HandleFunc("POST /{$}", func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(400)
			return
		}
		w.Write(body)
	})
}
