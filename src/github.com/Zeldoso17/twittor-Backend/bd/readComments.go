package bd

import (
	"context"
	"log"
	"time"

	"github.com/Zeldoso17/twittor-Backend/models"
	"go.mongodb.org/mongo-driver/bson"
)

func ReadComments(idtweet string) ([]*models.ReturnComments, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("comentarios")

	var results []*models.ReturnComments // We create a slice of pointers to ReturnComments

	condition := bson.M{
		"tweetid": idtweet,
	}

	cursor, err := col.Find(ctx, condition)
	if err != nil {
		log.Fatal(err.Error())
		return results, false, err
	}

	for cursor.Next(context.TODO()) {
		var register models.ReturnComments
		err := cursor.Decode(&register) // Here we decode the tweet into the register variable
		if err != nil {
			return results, false, err
		}
		results = append(results, &register)
	}
	return results, true, nil
}