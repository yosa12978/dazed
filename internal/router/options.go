package router

import (
	"os"

	"github.com/yosa12978/dazed/internal/logging"
)

type routerOptions struct {
	logger logging.Logger
}

func newDefaultOptions() routerOptions {
	return routerOptions{
		logger: logging.NewLogger(os.Stdout, "INFO"),
	}
}

func newOptions(opts ...optionFunc) routerOptions {
	options := newDefaultOptions()
	for _, opt := range opts {
		opt(&options)
	}
	return options
}

type optionFunc func(o *routerOptions)

func WithLogger(logger logging.Logger) optionFunc {
	return func(o *routerOptions) {
		o.logger = logger
	}
}
