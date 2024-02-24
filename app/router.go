package app

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/ipkalid/order-api/handler"
	order_repo "github.com/ipkalid/order-api/repository/order"
)

// LoadRouter returns an http.Handler that handles the routing for the order API.
func (a *App) loadRouter() {
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(ContentTypeMiddleware("application/json"))

	router.Route("/order", a.loadOrderRoute)

	a.router = router

	// return router
}

// loadOrderRoute loads the order routes into the provided router.
// It registers the following routes:
// - POST /: Creates a new order
// - GET /: Lists all orders
// - GET /{id}: Retrieves an order by ID
// - PUT /{id}: Updates an order by ID
// - DELETE /{id}: Deletes an order by ID
func (a *App) loadOrderRoute(router chi.Router) {
	orderRouter := handler.NewOrder(order_repo.RedisRepo{Client: a.client})
	router.Post("/", orderRouter.Create)
	router.Get("/", orderRouter.List)

	router.Get("/{id}", orderRouter.GetByID)
	router.Put("/{id}", orderRouter.UpdateByID)
	router.Delete("/{id}", orderRouter.DeleteByID)
}

func ContentTypeMiddleware(contentType string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Set Content-Type header
			w.Header().Set("Content-Type", contentType)
			// Call the next handler, which can be another middleware in the chain, or the final handler.
			next.ServeHTTP(w, r)
		})
	}
}
