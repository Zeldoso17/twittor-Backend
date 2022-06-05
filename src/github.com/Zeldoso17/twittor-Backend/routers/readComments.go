package routers

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"

	"github.com/Zeldoso17/twittor-Backend/bd"
)

func ReadComments(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	IDTweet, err := vars["IDTweet"]
	if !err {
		http.Error(w, "El ID del tweet no es valido!!", http.StatusBadRequest)
		return
	}

	comments, status, _ := bd.ReadComments(IDTweet)
	if !status {
		http.Error(w, "Error al leer los comentarios", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(comments)
}