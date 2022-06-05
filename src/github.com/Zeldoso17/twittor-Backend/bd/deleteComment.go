package bd

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func DeleteComment(idComment string, idUser string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor") // Here we get the database
	col := db.Collection("comentarios")  // Here we get the collection

	objIDComment, _ := primitive.ObjectIDFromHex(idComment)

	condition := bson.M{
		"_id": objIDComment,
		"userid":  idUser,
	}

	_, err := col.DeleteOne(ctx, condition)

	return err
}