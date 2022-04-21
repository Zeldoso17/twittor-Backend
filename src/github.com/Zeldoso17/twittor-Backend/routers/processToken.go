package routers

import (
	"errors"
	"strings"

	"github.com/Zeldoso17/twittor-Backend/bd"
	"github.com/Zeldoso17/twittor-Backend/models"
	jwt "github.com/dgrijalva/jwt-go"
)

// Email is the value of the email that will use in whole EndPoints
var Email string

// IDUser is the ID returned by model User, that will use in whole EndPoints
var IDUser string

// ProcessToken is a function that allows to extract its values
func ProcessToken(tk string) (*models.Claim, bool, string, error) {
	privateKey := []byte("Cesun2022_cuatri9")
	claims := &models.Claim{} // This is the struct that will be used to extract the values from the token

	splitToken := strings.Split(tk, "Bearer") // We split the token to get the token value
	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("formato de token invalido")
	}

	tk = strings.TrimSpace(splitToken[1]) // We get the token without the word Bearer

	tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token) (interface{}, error) {
		return privateKey, nil
	})
	if err == nil {
		_, found, _ := bd.UserAlreadyExist(claims.Email)
		if found {
			Email = claims.Email
			IDUser = claims.ID.Hex()
		}
		return claims, found, IDUser, nil
	}
	if !tkn.Valid {
		return claims, false, string(""), errors.New("token invalido")
	}

	return claims, false, string(""), err
}
