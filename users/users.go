package Users

import (
	"context"

	"dings-go/database"

	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/mongo"

	"net/http"

	token "dings-go/functions"
)

type User struct {
	Username string `bson:"username"`
	Password string `bson:"password"`
}

func GetAllUsers() ([]map[string]interface{}, error) {
	//Conexion a la base de datos
	database.ConnectDatabase()

	// Obtener la colección users
	collection := database.Client.Database("dings-test").Collection("users")

	// Obtener todos los documentos de la colección users
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}

	// Procesar los documentos y guardarlos en un slice de maps
	var results []map[string]interface{}
	for cursor.Next(context.Background()) {
		var result map[string]interface{}
		if err := cursor.Decode(&result); err != nil {
			return nil, err
		}
		results = append(results, result)
	}

	return results, nil
}

// Función para verificar si la contraseña y el usuario son correctos
func GetUsersByUsernameAndPassword(w http.ResponseWriter, r *http.Request, username string, password string) (map[string]interface{}, error) {
	// Conexion a la base de datos
	database.ConnectDatabase()

	// Obtener la colección users
	collection := database.Client.Database("dings-test").Collection("users")

	// Crear un filtro para buscar al usuario
	filter := bson.M{"username": username, "password": password}

	// Buscar al usuario en la colección
	var user User
	err := collection.FindOne(context.Background(), filter).Decode(&user)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			// El usuario no fue encontrado
			return map[string]interface{}{"authLogin": false}, nil
		} else {
			// Ocurrió un error al buscar el usuario
			return nil, err
		}
	}

	// Generar el token de autenticación en base64
	tokenAuth, err := token.GenerateAuthToken()

	token.CreateAuthTokenSession(w, r, tokenAuth)

	// El usuario fue encontrado, devolver el token de autenticación y el estado de autenticación como un mapa
	return map[string]interface{}{"authLogin": true, "token": tokenAuth, "username": user.Username}, nil
}
