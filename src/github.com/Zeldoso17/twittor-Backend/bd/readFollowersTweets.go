package bd

import (
	"context"
	"time"
	"fmt"

	"github.com/Zeldoso17/twittor-Backend/models"
	"go.mongodb.org/mongo-driver/bson"
)

// ReadAllUsers is a function that allows to read the tweets of all my followers
func ReadFollowersTweet(ID string, pagina int) ([]*models.ReturnFollowersTweets, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("relacion")

	skip := (pagina - 1) * 20 // Here we set the number of documents to be skipped

	conditions := make([]bson.M, 0) // Here we create a slice of bson.M
	conditions = append(conditions, bson.M{"$match": bson.M{"usuarioid": ID}}) // Here we add the condition to the slice
	conditions = append(conditions, bson.M{ // Here we are using the $lookup operator to join the two collections
		"$lookup": bson.M{ // Here we create the $lookup operator
			"from":	 "tweet", // Here we set the collection to be joined
			"localField": "usuariorelacionid", // Here we set the field of the current collection
			"foreignField": "userid", // Here we set the field of the collection to be joined
			"as": "tweet", // Here we set the name of the new collection
		}})
	
	conditions = append(conditions, bson.M{"$unwind": "$tweet"}) // Here we are using the $unwind operator to separate the documents
	conditions = append(conditions, bson.M{"$sort": bson.M{"tweet.fecha": -1}}) // Here we are using the $sort operator to sort the documents
	conditions = append(conditions, bson.M{"$skip": skip}) // Here we are using the $skip operator to skip the documents
	conditions = append(conditions, bson.M{"$limit": 20}) // Here we are using the $limit operator to limit the documents

	cursor, _ := col.Aggregate(ctx, conditions) // Here we are using the Aggregate function to execute the query
	var results []*models.ReturnFollowersTweets
	err := cursor.All(ctx, &results) // Here we are using the All function to decode the documents in the cursor into the results variable
	if err != nil {
		fmt.Println(err.Error())
		return results, false
	}
	return results, true
}