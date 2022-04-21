package routers

import (
	"encoding/json"
	"net/http"
	"time"
	"github.com/Zeldoso17/twittor-Backend/bd"
	"github.com/Zeldoso17/twittor-Backend/models"
)

// CreateTweet allows to create a new tweet in the database
func CreateTweet(w http.ResponseWriter, r *http.Request) {
	var mensaje models.Tweet
	err := json.NewDecoder(r.Body).Decode(&mensaje)

	if err != nil {
		http.Error(w, "Datos Incorrectos "+err.Error(), 400)
		return
	}

	registro := models.CreateTweet{ // Here we create a new struct to store the data
		UserID:  IDUser,
		Mensaje: mensaje.Mensaje,
		Fecha: time.Now(),
	}

	_, status, err := bd.InsertTweet(registro)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar insertar el registro. Reintente nuevamente " + err.Error(), http.StatusBadRequest)
		return
	}

	if !status {
		http.Error(w, "No se ha logrado insertar el registro", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}