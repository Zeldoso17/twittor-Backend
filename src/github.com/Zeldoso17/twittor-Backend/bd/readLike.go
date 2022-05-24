package bd

import (
	"context"
	"log"
	"time"

	"github.com/Zeldoso17/twittor-Backend/models"
	"go.mongodb.org/mongo-driver/bson"
)

func ReadLike(UserID string, TweetID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("likes")

	condition := bson.M{
		"userid": UserID,
		"tweetid": TweetID,
	}

	var result models.ResponseGiveLike
	err := col.FindOne(ctx, condition).Decode(&result)
	if err != nil {
		log.Println("Like no encontrado")
		return false, err
	}
	return true, nil
}