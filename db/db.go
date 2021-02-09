package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoClient recibe valor de la funcion DBConnect
var MongoClient = Connect()
var clientOptions = options.Client().ApplyURI("mongodb+srv://go_user_01:Acuario1936@cluster0.uwcgg.mongodb.net/test-api-go?retryWrites=true&w=majority")

// Connect maneja la Conexion a la base de MongoDB
func Connect() *mongo.Client {
	mongoClient, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return mongoClient
	}
	err = mongoClient.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return mongoClient
	}
	log.Println("Base de datos conectada")
	return mongoClient
}

// CheckConnection checkea que la conxión a la base de datos sea correcta
func CheckConnection() bool {
	err := MongoClient.Ping(context.TODO(), nil)
	if err != nil {
		return false
	}
	return true
}