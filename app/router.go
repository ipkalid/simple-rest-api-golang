package app

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/ipkalid/order-api/handler"
)

func LoadRouter() http.Handler {
	router := chi.NewRouter()

	router.Use(middleware.Logger)

	router.Route("/order", loadOrderRoute)

	return router
}

func loadOrderRoute(router chi.Router) {
	orderRouter := &handler.Order{}
	router.Post("/ ", orderRouter.Create)
	router.Get("/", orderRouter.List)

	router.Get("/{id}", orderRouter.GetByID)
	router.Put("/{id}", orderRouter.UpdateByID)
	router.Delete("/{id}", orderRouter.DeleteByID)
}
