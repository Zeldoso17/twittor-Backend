package bd

import (
	"context"
	"log"
	"time"

	"github.com/Zeldoso17/twittor-Backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ReadTweets allows to read whole tweets of the database
func ReadTweets(ID string, pagina int64) ([]*models.ReturnTweets, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("tweet")

	var results []*models.ReturnTweets // We create a slice of pointers to ReturnTweets

	condition := bson.M{
		"userid": ID,
	}

	options := options.Find()
	options.SetLimit(20)                               // Here we set the limit of tweets to be read
	options.SetSort(bson.D{{Key: "fecha", Value: -1}}) // Here we sort the tweets by date in descending order
	options.SetSkip((pagina - 1) * 20)                 // Here we skip the first 20 tweets for every page

	cursor, err := col.Find(ctx, condition, options)
	if err != nil {
		log.Fatal(err.Error())
		return results, false
	}

	for cursor.Next(context.TODO()) {
		var register models.ReturnTweets
		err := cursor.Decode(&register) // Here we decode the tweet into the register variable
		if err != nil {
			return results, false
		}
		results = append(results, &register)
	}
	return results, true
}
