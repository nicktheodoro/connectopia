package routers

import (
	"connectopia-api/src/controllers"
	"net/http"
)

var loginRoute = Route{
	URI:         "/login",
	Method:      http.MethodPost,
	HandlerFunc: controllers.Login,
	RequireAuth: false,
}
