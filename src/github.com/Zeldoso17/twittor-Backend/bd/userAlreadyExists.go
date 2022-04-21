package bd

import (
	"context"
	"time"

	"github.com/Zeldoso17/twittor-Backend/models"
	"go.mongodb.org/mongo-driver/bson"
)

// UserAlreadyExists is a function to check if a user already exists in the database
func UserAlreadyExist(email string) (models.Usuario, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor") // We are using the twittor database
	col := db.Collection("usuarios")  // We are using the usuarios collection

	condition := bson.M{"email": email} // We are creating a condition to find the user

	var result models.Usuario // We are creating a variable to store the user

	err := col.FindOne(ctx, condition).Decode(&result) // We are searching for the user in the database
	ID := result.ID.Hex()
	if err != nil {
		return result, false, ID
	}
	return result, true, ID
}
