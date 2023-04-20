package main

import (
	"log"
	"net/http"

	"dings-go/routes"
)

func main() {
	routes.SetupRoutes()

	log.Println("Iniciando servidor en http://localhost:8090...")
	log.Fatal(http.ListenAndServe(":8090", nil))
}
