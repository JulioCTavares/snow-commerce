package main

import (
	"db-project/internal/db"
	"db-project/internal/handler"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	db.InitDB()

	r := mux.NewRouter()
	r.HandleFunc("/users", handler.CreateUser).Methods("POST")
	r.HandleFunc("/users", handler.GetUsers).Methods("GET")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server running at http://localhost:%s/", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
