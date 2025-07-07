package main

import (
	"log"
	"net/http"

	"github.com/karahalil/backend-project/config"
	"github.com/karahalil/backend-project/db"
	"github.com/karahalil/backend-project/routes"
)

func main() {
	config.Load()
	db.Connect()
	router := routes.Setup()

	//Create initial db data
	err := db.RunSQLFile("create-tables.sql")
	if err != nil {
		log.Fatal("Error running SQL file:", err)
	}

	log.Fatal(http.ListenAndServe(":8080", router)) // Starts the server on port 8000
}
