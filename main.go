package main

import (
	"log"
	"net/http"

	"dings-go/routes"
)

func main() {
	routes.SetupRoutes()

	log.Println("Iniciando servidor en http://localhost:8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
