package routers

import (
	"connectopia-api/src/controllers"
	"net/http"
)

var userRoutes = []Route{
	{
		URI:         "/users",
		Method:      http.MethodGet,
		HandlerFunc: controllers.FindByNameOrUsername,
		RequireAuth: true,
	},
	{
		URI:         "/users/{id}",
		Method:      http.MethodGet,
		HandlerFunc: controllers.GetUserByID,
		RequireAuth: true,
	},
	{
		URI:         "/users",
		Method:      http.MethodPost,
		HandlerFunc: controllers.CreateUser,
		RequireAuth: false,
	},
	{
		URI:         "/users/{id}",
		Method:      http.MethodPut,
		HandlerFunc: controllers.UpdateUser,
		RequireAuth: true,
	},
	{
		URI:         "/users/{id}",
		Method:      http.MethodDelete,
		HandlerFunc: controllers.DeleteUser,
		RequireAuth: true,
	},
}
