package bd

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// DeleteTweet allows to delete a tweet from the database
func DeleteTweet(ID string, UserID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("tweet")

	objID, _ := primitive.ObjectIDFromHex(ID)

	condition := bson.M{ // Here we create a condition to delete the tweet
		"_id":    objID,
		"userid": UserID,
	}

	_, err := col.DeleteOne(ctx, condition) // Here we delete the tweet
	return err
}
