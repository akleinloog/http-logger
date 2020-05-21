package app

import (
	"github.com/akleinloog/http-logger/config"
	"github.com/akleinloog/http-logger/util/logger"
)

// App represents the overall application.
type App struct {
	logger *logger.Logger
	config *config.Config
}

// Instance creates a new App with config and logger..
func Instance() *App {
	config := config.AppConfig()
	logger := logger.New(config)
	return &App{logger: logger, config: config}
}

// Logger provides access to the global logger.
func (app *App) Logger() *logger.Logger {
	return app.logger
}

// Config provides access to the global configuration.
func (app *App) Config() *config.Config {
	return app.config
}
