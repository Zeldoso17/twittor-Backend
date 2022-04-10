package routers

import (
	"encoding/json"
	"net/http"
	"github.com/Zeldoso17/twittor-Backend/src/github.com/Zeldoso17/twittor-Backend/bd"
	"github.com/Zeldoso17/twittor-Backend/src/github.com/Zeldoso17/twittor-Backend/models"
)

// Register is a function to create a new user in the database
func Register(w http.ResponseWriter, r *http.Request){
	var t models.Usuario
	err := json.NewDecoder(r.Body).Decode(&t) // We're reading the body of the request
	if err != nil {
		http.Error(w, "Something went wrong with the request body" + err.Error(), 400)
		return
	}
	if len(t.Email) == 0 {
		
	}
}