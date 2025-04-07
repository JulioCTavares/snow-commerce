package main

import (
	"db-project/internal/db"
	"db-project/internal/routes"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	db.InitDB()

	r := routes.SetupRouter()

	printRoutes(r) // 👉 Imprime todas as rotas

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server running at http://localhost:%s/", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}

func printRoutes(r *mux.Router) {
	fmt.Println("\n📋 Rotas disponíveis:")

	err := r.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
		path, err := route.GetPathTemplate()
		if err != nil {
			return err
		}
		methods, err := route.GetMethods()
		if err != nil {
			return nil
		}
		fmt.Printf("👉 %s %s\n", methods, path)
		return nil
	})

	if err != nil {
		log.Println("Erro ao listar rotas:", err)
	}
}
