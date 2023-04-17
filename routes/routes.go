package routes

import (
	Users "dings-go/users"
	"encoding/json"
	"log"
	"net/http"
)

func UsersManagement(w http.ResponseWriter, r *http.Request) {
	// Lógica para la ruta 1
	users, err := Users.GetAllUsers()
	if err != nil {
		log.Fatal(err)
	} else {
		json.NewEncoder(w).Encode(users)
	}

}

func Handler2(w http.ResponseWriter, r *http.Request) {
	// Lógica para la ruta 2
}

func SetupRoutes() {
	http.HandleFunc("/users", UsersManagement)
	http.HandleFunc("/ruta2", Handler2)
}
