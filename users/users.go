package Users

import (
	"context"

	"dings-go/database"

	"go.mongodb.org/mongo-driver/bson"
)

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
