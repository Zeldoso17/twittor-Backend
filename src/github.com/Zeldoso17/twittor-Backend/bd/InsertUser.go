package bd

import (
	"context"
	"time"

	"github.com/Zeldoso17/twittor-Backend/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// InsertUser is a function to create a new user in the database
func InsertUser(user models.Usuario) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second) // Here we create a context with a timeout of 15 seconds
	defer cancel()                                                           // here we close the context

	db := MongoCN.Database("twittor") // Here we get the database
	col := db.Collection("usuarios")  // Here we get the collection

	user.Password, _ = EncryptPassword(user.Password) // Here we encrypt the password

	result, err := col.InsertOne(ctx, user) // Here we insert the user in the database
	if err != nil {
		return "", false, err
	}

	ObjID, _ := result.InsertedID.(primitive.ObjectID) // Here we get the ID of the user inserted
	return ObjID.String(), true, nil                   // Here we return the ID of the user inserted

}
