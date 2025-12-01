package main

import (
	"log"
	"os"

	"question-answer/database"

	"github.com/pressly/goose/v3"
)

func main() {
	if err := database.Init(); err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	args := os.Args[1:]
	if len(args) < 1 {
		log.Fatal("Usage: migrate [up|down|status|create] [args]")
	}

	command := args[0]
	dir := "migrations"

	db, err := database.GetDB().DB()
	if err != nil {
		log.Fatal("Failed to get database connection:", err)
	}

	if err := goose.Run(command, db, dir, args[1:]...); err != nil {
		log.Fatal("Migration failed:", err)
	}

	log.Println("Migration command completed:", command)
}
