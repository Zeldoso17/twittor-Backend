package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/Zeldoso17/twittor-Backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ReadAllUsers is a function that allows to read all users registered
func ReadAllUsers(ID string, page int64, search string, tipo string) ([]*models.Usuario, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("usuarios")

	var results []*models.Usuario

	findOptions := options.Find()
	findOptions.SetSkip((page - 1) * 20) // Here we set the number of documents to be skipped
	findOptions.SetLimit(20)             // Here we set the limit of the number of documents to be returned

	query := bson.M{
		"nombre": bson.M{"$regex": `(?i)` + search}, // Here we create a condition to be used in the query
	}

	cur, err := col.Find(ctx, query, findOptions) // Here we execute the query and save the results in a cursor
	if err != nil {
		fmt.Println(err.Error())
		return results, false
	}

	var found, include bool

	for cur.Next(ctx) {
		var user models.Usuario  // Here we create a variable of type models.Usuario
		err := cur.Decode(&user) // Here we decode the document in the cursor into the user variable
		if err != nil {
			fmt.Println(err.Error())
			return results, false
		}

		var rela models.Relation
		rela.UserID = ID                       // Here we set the user ID to be compared with the user ID of the user that is logged in
		rela.UsuarioRelacionID = user.ID.Hex() // Here we set the user ID of the user that we want to compare with the user ID of the user that is logged in

		include = false

		found, err = ReadRelation(rela) // Here we check if the user that we want to compare with the user that is logged in is in the relation
		fmt.Println(err)

		if tipo == "new" && !found {
			include = true
		}
		if tipo == "follow" && found {
			include = true
		}

		if rela.UsuarioRelacionID == ID {
			include = false
		}

		if include {
			user.Password = ""  // Here we erase the password of the user
			user.Biografia = "" // Here we erase the biography of the user
			user.SitioWeb = ""  // Here we erase the web site of the user
			user.Ubicacion = "" // Here we erase the location of the user
			user.Banner = ""    // Here we erase the banner of the user
			user.Email = ""     // Here we erase the email of the user

			results = append(results, &user) // Here we add the user to the results
		}
	}

	err = cur.Err()
	if err != nil {
		fmt.Println(err.Error())
		return results, false
	}
	cur.Close(ctx)
	return results, true
}