package main

import (
	"log"
	"question-answer/database"
	"question-answer/router"
	"question-answer/routes"
)

func main() {
	if err := database.Init(); err != nil {
		log.Fatal("Failed to initialize database:", err)
	}
	allRoutes := routes.RegisterAllRoutes()

	appRouter := router.New(allRoutes)

	server := NewApi(appRouter, "8080")
	log.Fatal(server.Start())
}
