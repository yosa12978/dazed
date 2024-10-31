package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/yosa12978/dazed/internal/config"
)

func main() {
	conf := config.Get()

	http.HandleFunc("GET /{$}", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello dazed")
	})
	http.HandleFunc("GET /config", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "%+v", conf)
	})
	http.HandleFunc("POST /", func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(400)
			return
		}
		w.Write(body)
	})
	http.ListenAndServe(conf.Server.Addr, nil)
}
