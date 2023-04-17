package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

func ConnectDatabase() {
	// Configurar opciones de conexión
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Conectar a la base de datos
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Comprobar la conexión
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	// Asignar el cliente de MongoDB a la variable global
	Client = client
}
