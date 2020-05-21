package app

import (
	"github.com/akleinloog/http-logger/config"
	"github.com/akleinloog/http-logger/router"
	"github.com/akleinloog/http-logger/util/logger"
	"github.com/go-chi/chi"
)

// Server represents the overall application.
type Server struct {
	logger *logger.Logger
	config *config.Config
	router *chi.Mux
}

// Instance creates a new Server with config and logger..
func Instance() *Server {
	return &Server{logger: logger.New(config.AppConfig()), config: config.AppConfig(), router: router.New()}
}

// Logger provides access to the global logger.
func (app *Server) Logger() *logger.Logger {
	return app.logger
}

// Config provides access to the global configuration.
func (app *Server) Config() *config.Config {
	return app.config
}

// Config provides access to the global configuration.
func (app *Server) Router() *chi.Mux {
	return app.router
}
