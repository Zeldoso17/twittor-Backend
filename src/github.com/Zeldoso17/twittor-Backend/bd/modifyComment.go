package bd

import (
	"context"
	"time"

	"github.com/Zeldoso17/twittor-Backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ModifyComment(c models.CreateComment, CommentID string, UserID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("comentarios")

	comment := validationscomment(c)

	updtString := bson.M{
		"$set": comment,
	}

	objID, _ := primitive.ObjectIDFromHex(CommentID)
	filter := bson.M{"_id": bson.M{"$eq": objID}, "userid": bson.M{"$eq": UserID}}

	_, err := col.UpdateOne(ctx, filter, updtString)
	if err != nil {
		return false, err
	}
	return true, nil
}

func validationscomment(c models.CreateComment) map[string]interface{} {
	registerValidations := make(map[string]interface{})

	if len(c.UserID) > 0 {
		registerValidations["userid"] = c.UserID
	}
	if len(c.Mensaje) > 0 {
		registerValidations["mensaje"] = c.Mensaje
	}
	registerValidations["fecha"] = c.Fecha
	if len(c.TweetID) > 0 {
		registerValidations["tweetid"] = c.TweetID
	}

	return registerValidations
}