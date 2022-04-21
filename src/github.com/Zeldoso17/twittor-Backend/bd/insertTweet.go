package bd

import (
	"context"
	"time"

	"github.com/Zeldoso17/twittor-Backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InsertTweet(t models.CreateTweet) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second) // Here we create a context with a timeout of 15 seconds
	defer cancel()                                                           // here we close the context

	db := MongoCN.Database("twittor") // Here we get the database
	col := db.Collection("tweet")     // Here we get the collection

	registro := bson.M{
		"userid":  t.UserID,
		"mensaje": t.Mensaje,
		"fecha":   t.Fecha,
	}
	result, err := col.InsertOne(ctx, registro) // Here we insert the tweet
	if err != nil {
		return "", false, err
	}
	objID, _ := result.InsertedID.(primitive.ObjectID) // Here we get the object ID
	return objID.String(), true, nil
}
