package routers

import (
	"connectopia-api/src/controllers"
	"net/http"
)

var publicationRoutes = []Route{
	{
		URI:         "/publications",
		Method:      http.MethodPost,
		HandlerFunc: controllers.CreatePublication,
		RequireAuth: true,
	},
	{
		URI:         "/publications",
		Method:      http.MethodGet,
		HandlerFunc: controllers.GetPublications,
		RequireAuth: true,
	},
	{
		URI:         "/publications/{id}",
		Method:      http.MethodGet,
		HandlerFunc: controllers.GetPublication,
		RequireAuth: true,
	},
	{
		URI:         "/publications/{id}",
		Method:      http.MethodPut,
		HandlerFunc: controllers.UpdatePublication,
		RequireAuth: true,
	},
	{
		URI:         "/publications/{id}",
		Method:      http.MethodDelete,
		HandlerFunc: controllers.DeletePublication,
		RequireAuth: true,
	},
	{
		URI:         "/users/{id}/publications",
		Method:      http.MethodGet,
		HandlerFunc: controllers.GetPublicationsByUser,
		RequireAuth: true,
	},
	{
		URI:         "/publications/{id}/like",
		Method:      http.MethodPost,
		HandlerFunc: controllers.LikePublication,
		RequireAuth: true,
	},
	{
		URI:         "/publications/{id}/unlike",
		Method:      http.MethodPost,
		HandlerFunc: controllers.UnlikePublication,
		RequireAuth: true,
	},
}
