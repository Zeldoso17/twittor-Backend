package bd

import (
	"context"
	"time"
	"fmt"
	"github.com/Zeldoso17/twittor-Backend/models"
	"go.mongodb.org/mongo-driver/bson"
)

// ReadRelation is a function that reads a relation between two users
func ReadRelation(t models.Relation) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("relacion")

	condition := bson.M{ // Here I'm searching for the relation
		"usuarioid":         t.UserID,
		"usuariorelacionid": t.UsuarioRelacionID,
	}

	var result models.Relation
	fmt.Println(result)
	err := col.FindOne(ctx, condition).Decode(&result)
	if err != nil {
		fmt.Println(err.Error())
		return false, err
	}
	return true, nil
}