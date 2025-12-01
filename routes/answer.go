package routes

import (
	"net/http"
)

type AnswerHandler interface {
	CreateAnswer(w http.ResponseWriter, r *http.Request)
	GetAnswer(w http.ResponseWriter, r *http.Request)
	DeleteAnswer(w http.ResponseWriter, r *http.Request)
}

func RegisterAnswerRoutes(h AnswerHandler) []RouteDefinition {
	return []RouteDefinition{
		{
			Method:  "POST",
			Path:    "/questions/{id}/answers/",
			Handler: h.CreateAnswer,
		},
		{
			Method:  "GET",
			Path:    "/answers/{id}",
			Handler: h.GetAnswer,
		},
		{
			Method:  "DELETE",
			Path:    "/answers/{id}",
			Handler: h.DeleteAnswer,
		},
	}
}
