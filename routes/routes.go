package routes

import (
	"net/http"
	"question-answer/handlers"
)

type RouteDefinition struct {
	Method  string
	Path    string
	Handler http.HandlerFunc
}

func RegisterAllRoutes() []RouteDefinition {
	var routes []RouteDefinition

	questionHandler := handlers.NewQuestionHandler(nil)
	questionRoutes := RegisterQuestionRoutes(questionHandler)
	routes = append(routes, questionRoutes...)

	answerHandler := handlers.NewAnswerHandler(nil, nil, nil)
	answerRoutes := RegisterAnswerRoutes(answerHandler)
	routes = append(routes, answerRoutes...)

	userHandler := handlers.NewUserHandler()
	userRoutes := RegisterUserRoutes(userHandler)
	routes = append(routes, userRoutes...)

	return routes
}
