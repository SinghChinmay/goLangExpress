package main

import (
	"goLangExpress/database"
	"goLangExpress/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	database.InitDB()

	r := mux.NewRouter()
	routes.SetupRoutes(r)

	log.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}