package routers

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Route represents an API route.
type Route struct {
	URI         string
	Method      string
	HandlerFunc http.HandlerFunc
	RequireAuth bool
}

// Configure adds all routes to the router.
func Configure(router *mux.Router) *mux.Router {
	routes := userRoutes

	for _, route := range routes {
		router.HandleFunc(route.URI, route.HandlerFunc).Methods(route.Method)
	}

	return router
}
