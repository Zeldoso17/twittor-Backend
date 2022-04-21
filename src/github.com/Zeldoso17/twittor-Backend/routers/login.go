package routers

import (
	"encoding/json"
	"net/http"
	"github.com/Zeldoso17/twittor-Backend/bd"
	"github.com/Zeldoso17/twittor-Backend/jwt"
	"github.com/Zeldoso17/twittor-Backend/models"
)
// Login is a function that handles the login process
func Login(w http.ResponseWriter, r *http.Request){
	w.Header().Add("content-type", "application/json") // Here we are adding the content-type header

	var t models.Usuario // Here we are creating a variable to store the user

	err := json.NewDecoder(r.Body).Decode(&t) // Here we are decoding the body of the request into the variable t
	if err != nil {
		http.Error(w, "Usuario y/o Contraseña incorrectos" + err.Error(), 400)
		return
	}

	// We start with validations
	if len(t.Email) == 0 {
		http.Error(w, "El email es requerido", 400)
		return
	}
	document, exist := bd.LoginTry(t.Email, t.Password)
	if !exist {
		http.Error(w, "Usuario y/o Contraseña incorrectos", 400)
		return
	}

	// Here we are creating the token
	key, err := jwt.GenerateJWT(document)
	if err != nil {
		http.Error(w, "Ocurrió un error al generar el token" + err.Error(), 400)
		return
	}

	resp := models.ResponseLogin {
		Token: key,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
	
	/* Here we are setting the expiration time of the token
	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name: "token",
		Value: key,
		Expires: expirationTime,
	}) */
}