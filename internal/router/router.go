package router

import (
	"fmt"
	"io"
	"net/http"
)

func NewRouter(opts ...optionFunc) http.Handler {
	options := newOptions(opts...)
	router := http.NewServeMux()
	addV1Routes(router, options)
	return router
}

func addV1Routes(router *http.ServeMux, options routerOptions) {
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
