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
	{
		URI:         "/users/{id}/follow",
		Method:      http.MethodPost,
		HandlerFunc: controllers.FollowUser,
		RequireAuth: true,
	},
	{
		URI:         "/users/{id}/unfollow",
		Method:      http.MethodPost,
		HandlerFunc: controllers.UnfollowUser,
		RequireAuth: true,
	},
}
