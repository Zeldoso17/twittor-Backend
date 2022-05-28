package bd

import (
    "context"
	"log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/* MongoCN is a DataBase connection Object */
var MongoCN = ConnectionBD()
var clientOptions = options.Client().ApplyURI("mongodb+srv://Zeldoso17:twittoroso17@twittor.emna0.mongodb.net/myFirstDatabase?retryWrites=true&w=majority")

/* ConnectionBD is a function to allows me to connect to the DataBase */
func ConnectionBD() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	
	if err != nil{
		log.Fatal(err.Error())
		return client // retorna el objeto cliente
	}

	err = client.Ping(context.TODO(), nil) // Revisa si la conexion esta activa

	if err != nil{
		log.Fatal(err.Error())
		return client
	}

	log.Println("Conexión exitosa con la BD")
	return client
}

/* ConnectionStatus is the function to allows me to check the Connection Status */
func ConnectionStatus() bool {
	err := MongoCN.Ping(context.TODO(), nil)

	return err == nil
}