package routers

import (
	"encoding/json"
	"net/http"
	"github.com/Zeldoso17/twittor-Backend/bd"
	"github.com/Zeldoso17/twittor-Backend/models"
)

// Register is a function to create a new user in the database
func Register(w http.ResponseWriter, r *http.Request){
	var t models.Usuario
	err := json.NewDecoder(r.Body).Decode(&t) // We're reading the body of the request
	if err != nil {
		http.Error(w, "Algo fué mal con los datos del request" + err.Error(), 400)
		return
	}

	validations(t, w)
	
	w.WriteHeader(http.StatusCreated)
}

func validations(user models.Usuario, w http.ResponseWriter){
	if len(user.Email) == 0 {
		http.Error(w, "Email is required", 400)
		return
	}
	if len(user.Password) < 6 {
		http.Error(w, "Password must be at least 6 characters", 400)
		return
	}
	_, encontrado, _ := bd.UserAlreadyExist(user.Email)
	if encontrado {
		http.Error(w, "User already exists", 400)
		return
	}

	_, status, err := bd.InsertUser(user)
	if err != nil {
		http.Error(w, "Something went wrong with Register User" + err.Error(), 400)
		return
	}
	if !status {
		http.Error(w, "User wasn't registered" + err.Error(), 400)
		return
	}
}