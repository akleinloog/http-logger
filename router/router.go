package router

import (
	"net/http"

	"github.com/akleinloog/http-logger/util/logger"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

// New instantiates a new router.
func New() *chi.Mux {

	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(logger.RequestLogger)

	router.Get("/", handleIndex)

	return router
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
}
