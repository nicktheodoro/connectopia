package routers

import (
	"connectopia-api/src/controllers"
	"net/http"
)

var userRoutes = []Route{
	{
		URI:         "/users",
		Method:      http.MethodGet,
		HandlerFunc: controllers.GetUsers,
		RequireAuth: false,
	},
	{
		URI:         "/users/{id}",
		Method:      http.MethodGet,
		HandlerFunc: controllers.GetUserByID,
		RequireAuth: false,
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
		RequireAuth: false,
	},
	{
		URI:         "/users/{id}",
		Method:      http.MethodDelete,
		HandlerFunc: controllers.DeleteUser,
		RequireAuth: false,
	},
}
