package main

import (
	"db-project/internal/db"
	"db-project/internal/routes"
	"log"
	"net/http"
	"os"
)

func main() {
	db.InitDB()

	r := routes.SetupRouter()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server running at http://localhost:%s/", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
