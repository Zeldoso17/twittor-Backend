package routers

import (
	"encoding/json"
	"net/http"
	"github.com/Zeldoso17/twittor-Backend/bd"
	"github.com/Zeldoso17/twittor-Backend/models"
)

// ModifyProfile is a function that modify a User Profile
func ModifyProfile(w http.ResponseWriter, r *http.Request){
	var t models.Usuario // Here I'm creating a variable of type Usuario
	err := json.NewDecoder(r.Body).Decode(&t) // Here I'm decoding the body of the request to the variable t
	if err != nil {
		http.Error(w, "Datos Incorrectos"+err.Error(), 400)
		return
	}

	var status bool
	
	status, err = bd.ModifyRegister(t, IDUser) // Here I'm calling the ModifyRegister function to modify the register
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar modificar el registro. Reintente nuevamente "+err.Error(), 400)
		return
	}
	if !status {
		http.Error(w, "No se ha logrado modificar el registro del usuario", 400)
		return
	}
	w.WriteHeader(http.StatusCreated)
}