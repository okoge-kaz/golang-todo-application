package main

import (
	"fmt"
	"os"

	"github.com/okoge-kaz/golang-todo-application/server/db"
	"github.com/okoge-kaz/golang-todo-application/server/router"
)

func main() {
	// initialize database connection
	db.Init(
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	// initialize router
	router := router.Init()

	// start server
	const port = 8000
	router.Run(fmt.Sprintf(":%d", port))
}
