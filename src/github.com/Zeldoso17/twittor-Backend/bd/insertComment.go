package bd

import (
	"context"
	"time"

	"github.com/Zeldoso17/twittor-Backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InsertComment(c models.CreateComment) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second) // Here we create a context with a timeout of 15 seconds
	defer cancel()                                                           // here we close the context

	db := MongoCN.Database("twittor") // Here we get the database
	col := db.Collection("comentarios")  // Here we get the collection

	register := bson.M {
		"userid":  c.UserID,
		"mensaje": c.Mensaje,
		"fecha":   c.Fecha,
		"tweetid": c.TweetID,
	}

	result, err := col.InsertOne(ctx, register) // Here we insert the tweet
	if err != nil {
		return "", false, err
	}

	objID, _ := result.InsertedID.(primitive.ObjectID) // Here we get the object ID
	return objID.String(), true, nil
}

