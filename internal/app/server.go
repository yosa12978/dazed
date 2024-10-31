package app

import (
	"net/http"

	"github.com/yosa12978/dazed/internal/logging"
	"github.com/yosa12978/dazed/internal/router"
)

func newServer(addr string, logger logging.Logger) http.Server {
	router := router.NewRouter(router.WithLogger(logger))
	return http.Server{
		Addr:    addr,
		Handler: router,
	}
}
