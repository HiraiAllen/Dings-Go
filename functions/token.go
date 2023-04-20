package token

import (
	"crypto/rand"
	"encoding/base64"

	"net/http"

	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("secret-key"))

// Función para generar un token de autenticación en base64
func GenerateAuthToken() (string, error) {
	// Generar un número aleatorio de 32 bytes
	tokenBytes := make([]byte, 32)
	_, err := rand.Read(tokenBytes)
	if err != nil {
		return "", err
	}

	// Codificar el número aleatorio en base64
	token := base64.StdEncoding.EncodeToString(tokenBytes)

	return token, nil
}

// Función pública para crear la variable de sesión authToken
func CreateAuthTokenSession(w http.ResponseWriter, r *http.Request, token string) {
	// Crear una variable de sesión authToken
	session, err := store.Get(r, "authToken")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	session.Values["authToken"] = token

	// Guardar la variable de sesión en la tienda de sesiones
	err = session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func DeleteAuthTokenSession(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "authToken")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	session.Values["authToken"] = nil
	err = session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
