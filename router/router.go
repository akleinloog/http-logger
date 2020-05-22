package router

import (
	"context"
	"fmt"
	"github.com/akleinloog/http-logger/util/logger"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net/http"
)

// New instantiates a new router.
func New() *chi.Mux {

	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	//router.Use(middleware.Logger)
	router.Use(logger.RequestLogger)
	router.Use(middleware.Recoverer)
	router.Use(addUserContext)

	router.Get("/", handleIndex)
	router.Get("/hello", handleHello())

	return router
}

func addUserContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "user", "123")
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func handleIndex(w http.ResponseWriter, r *http.Request) {

	//w.Header().Set("Content-Length", "12")
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")

	w.Write([]byte("Request Received"))
}

func handleHello() http.HandlerFunc {
	greeting := "Hello"
	return func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value("user").(string)
		w.Write([]byte(fmt.Sprintf("%s %s", greeting, user)))
	}
}
