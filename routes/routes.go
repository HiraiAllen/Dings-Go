package routes

import (
	Users "dings-go/users"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func UsersManagement(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*") // Cambiar el asterisco por el dominio que necesites permitir
	// Lógica para la ruta 1
	users, err := Users.GetAllUsers()
	if err != nil {
		log.Fatal(err)
	} else {
		json.NewEncoder(w).Encode(users)
	}

}

func UsersLoginConfirm(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With")
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Error al analizar los datos del formulario", http.StatusBadRequest)
			return
		}

		username := r.Form.Get("username")
		password := r.Form.Get("password")
		fmt.Printf("Nombre de usuario: %s, Contraseña: %s\n", username, password)

		confirmData, err := Users.GetUsersByUsernameAndPassword(w, r, username, password)
		//Convertir el map que llega a []byte
		jsonBytes, err := json.Marshal(confirmData)
		if err != nil {
			// manejar el error
			fmt.Printf("Error en el parseo de los datos")
		}

		w.WriteHeader(http.StatusOK)
		w.Write(jsonBytes)
	} else {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
	}
}

func SetupRoutes() {
	http.HandleFunc("/users", UsersManagement)
	http.HandleFunc("/login", UsersLoginConfirm)
}
