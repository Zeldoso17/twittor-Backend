package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/Zeldoso17/twittor-Backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ProfileSearch(ID string) (models.Usuario, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("usuarios")

	var profile models.Usuario
	objID, _ := primitive.ObjectIDFromHex(ID) // Here we convert the ID to a ObjectID

	condition := bson.M{"_id": objID,}

	err := col.FindOne(ctx, condition).Decode(&profile) // Here we search for the user in the database
	profile.Password = ""                               // We don't want to send the password to the frontend
	if err != nil {
		fmt.Println("Registro no encontrado " + err.Error())
		return profile, err
	}
	return profile, nil
}