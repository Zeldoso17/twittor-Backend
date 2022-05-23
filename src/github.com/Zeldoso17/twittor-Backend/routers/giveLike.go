package routers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/Zeldoso17/twittor-Backend/bd"
	"github.com/Zeldoso17/twittor-Backend/models"
)

func GiveLike(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	IDTweet, err := vars["IDTweet"]
	if !err{
		http.Error(w, "El ID del tweet no es valido!!", http.StatusBadRequest)
		return
	}

	register := models.GiveLike {
		UserID: IDUser,
		TweetID: IDTweet,
	}

	_, status, errlike := bd.GiveLike(register)

	if errlike != nil {
		http.Error(w, "Ocurrió un error al intentar dar like. Reintente nuevamente "+errlike.Error(), 400)
		return
	}

	if !status {
		http.Error(w, "No se ha logrado insertar el registro", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}