package app

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/yosa12978/dazed/internal/config"
	"github.com/yosa12978/dazed/internal/logging"
)

func Run() error {
	ctx, cancel := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)
	defer cancel()

	conf := config.Get()
	logger := logging.NewLogger(os.Stdout, conf.App.LogLevel)
	server := newServer(conf.Server.Addr, logger)

	doneCh := make(chan struct{}, 1)
	go func() {
		if err := server.ListenAndServe(); err != nil &&
			!errors.Is(err, http.ErrServerClosed) {
			logger.Error("server error", "error", err.Error())
		}
		close(doneCh)
	}()

	select {
	case <-doneCh:
	case <-ctx.Done():
		timeout, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		return server.Shutdown(timeout)
	}
	return nil
}
