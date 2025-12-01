package main

import (
	"log"
	"net/http"
	"question-answer/router"
)

type Server struct {
	router *router.Router
	port   string
}

func NewApi(router *router.Router, port string) *Server {
	return &Server{
		router: router,
		port:   port,
	}
}

func (s *Server) Start() error {
	addr := ":" + s.port

	log.Printf("Server is running on http://localhost%s", addr)
	return http.ListenAndServe(addr, s.router)
}
