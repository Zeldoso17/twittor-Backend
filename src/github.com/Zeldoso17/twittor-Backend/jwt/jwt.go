package jwt

import (
	"time"
	jwt "github.com/dgrijalva/jwt-go" // We create a alias for the package jwt-go
	"github.com/Zeldoso17/twittor-Backend/models"
)

// GenerateJWT is a function that generates the JWT
func GenerateJWT(t models.Usuario) (string, error) {
	privateKey := []byte("Cesun2022_cuatri9") // We create a key for the token
	payload := jwt.MapClaims{
		"email": t.Email,
		"nombre": t.Nombre,
		"apellidos": t.Apellidos,
		"fecha_nacimiento": t.FechaNacimiento,
		"biografia": t.Biografia,
		"ubicacion": t.Ubicacion,
		"sitioweb": t.SitioWeb,
		"_id": t.ID.Hex(),
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(privateKey)
	if err != nil {
		return tokenStr, err
	}
	return tokenStr, nil
}