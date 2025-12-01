package routes

import (
	"net/http"
)

type UserHandler interface {
	CreateUser(w http.ResponseWriter, r *http.Request)
}

func RegisterUserRoutes(h UserHandler) []RouteDefinition {
	return []RouteDefinition{
		{
			Method:  "GET",
			Path:    "/user",
			Handler: h.CreateUser,
		},
	}
}
