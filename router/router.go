package router

import (
	"context"
	"fmt"
	"github.com/akleinloog/http-logger/app"
	"github.com/akleinloog/http-logger/util/logger"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net/http"
)

// New initializes a new router.
func New(app *app.App) *chi.Mux {

	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	//router.Use(middleware.Logger)
	router.Use(logger.RequestLogger)
	router.Use(middleware.Recoverer)
	router.Use(addUserContext)
	//router.Get("/", func(w http.ResponseWriter, r *http.Request) {
	//	w.Write([]byte("welcome"))
	//})
	router.Get("/", handleIndex)

	//router.MethodFunc("GET", "/", HandleIndex)
	return router
}

func addUserContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "user", "123")
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func handleIndex(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Length", "12")
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")

	//w.Write([]byte("Welcome!"))

	user := r.Context().Value("user").(string)
	w.Write([]byte(fmt.Sprintf("Hi %s", user)))
}
