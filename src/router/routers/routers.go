package routers

import (
	"connectopia-api/src/middlewares"
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
func Configure(r *mux.Router) *mux.Router {
	routes := userRoutes
	routes = append(routes, loginRoute)

	for _, route := range routes {
		if route.RequireAuth {
			r.HandleFunc(route.URI, middlewares.Logger(middlewares.Authorize(route.HandlerFunc))).Methods(route.Method)
		} else {
			r.HandleFunc(route.URI, middlewares.Logger(route.HandlerFunc)).Methods(route.Method)
		}
	}

	return r
}
