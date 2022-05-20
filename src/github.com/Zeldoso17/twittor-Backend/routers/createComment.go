package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Zeldoso17/twittor-Backend/bd"
	"github.com/Zeldoso17/twittor-Backend/models"
	"github.com/gorilla/mux"
)

func CreateComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) // obtenemos el parametro id que esta en la url del tweet que se va a comentar

	IDTweet, err := vars["IDTweet"]
	if !err {
		http.Error(w, "El ID del tweet no es valido!!", http.StatusBadRequest)
		return
	}

	var comment models.Comment // creamos una variable de tipo comment para poder guardar los datos del comentario
	err2 := json.NewDecoder(r.Body).Decode(&comment) // decodificamos el json que viene en el body y lo guardamos en la variable comment

	if err2 != nil {
		http.Error(w, "Datos Incorrectos "+err2.Error(), 400)
		return
	}

	register := models.CreateComment { // creamos una variable de tipo createcomment para poder guardar los datos del comentario
		UserID: IDUser,
		Mensaje: comment.Mensaje,
		Fecha: time.Now(),
		TweetID: IDTweet,
	}
	
	_, status, errcomm := bd.InsertComment(register)
	if errcomm != nil {
		http.Error(w, "Ocurrió un error al intentar crear el comentario. Reintente nuevamente "+errcomm.Error(), 400)
		return
	}

	if !status {
		http.Error(w, "No se ha logrado insertar el registro", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}