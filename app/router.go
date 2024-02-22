package app

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func LoadRouter() http.Handler {
	router := chi.NewRouter()

	router.Use(middleware.Logger)

	router.Get("/", basic)

	return router
}

func basic(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	fmt.Fprintf(w, "hello world\n")
}
