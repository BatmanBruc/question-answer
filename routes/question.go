package routes

import (
	"net/http"
)

type QuestionHandler interface {
	GetQuestions(w http.ResponseWriter, r *http.Request)
	CreateQuestion(w http.ResponseWriter, r *http.Request)
	GetQuestion(w http.ResponseWriter, r *http.Request)
	DeleteQuestion(w http.ResponseWriter, r *http.Request)
}

func RegisterQuestionRoutes(h QuestionHandler) []RouteDefinition {
	return []RouteDefinition{
		{
			Method:  "GET",
			Path:    "/questions/",
			Handler: h.GetQuestions,
		},
		{
			Method:  "POST",
			Path:    "/questions/",
			Handler: h.CreateQuestion,
		},
		{
			Method:  "GET",
			Path:    "/questions/{id}",
			Handler: h.GetQuestion,
		},
		{
			Method:  "DELETE",
			Path:    "/questions/{id}",
			Handler: h.DeleteQuestion,
		},
	}
}
