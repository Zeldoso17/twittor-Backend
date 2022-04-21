package bd

import (
	"context"
	"time"

	"github.com/Zeldoso17/twittor-Backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ModifyRegister is a function that modify a User Profile
func ModifyRegister(u models.Usuario, ID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("usuarios")

	register := validations(u) // Here I'm calling the validations function to modify the register

	updtString := bson.M{
		"$set": register,
	}

	objID, _ := primitive.ObjectIDFromHex(ID)
	filter := bson.M{"_id": bson.M{"$eq": objID}}

	_, err := col.UpdateOne(ctx, filter, updtString)
	if err != nil {
		return false, err
	}
	return true, nil
}

// validations is a function that validates the register
func validations(usu models.Usuario) map[string]interface{} {
	registerValidations := make(map[string]interface{})

	if len(usu.Nombre) > 0 {
		registerValidations["nombre"] = usu.Nombre
	}
	if len(usu.Apellidos) > 0 {
		registerValidations["apellidos"] = usu.Apellidos
	}
	registerValidations["fechaNacimiento"] = usu.FechaNacimiento
	if len(usu.Avatar) > 0 {
		registerValidations["avatar"] = usu.Avatar
	}
	if len(usu.Banner) > 0 {
		registerValidations["banner"] = usu.Banner
	}
	if len(usu.Biografia) > 0 {
		registerValidations["biografia"] = usu.Biografia
	}
	if len(usu.Ubicacion) > 0 {
		registerValidations["ubicacion"] = usu.Ubicacion
	}
	if len(usu.SitioWeb) > 0 {
		registerValidations["sitioWeb"] = usu.SitioWeb
	}
	return registerValidations
}
