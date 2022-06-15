package routers

import (
	"net/http"
	"github.com/gorilla/mux"

	"github.com/Zeldoso17/twittor-Backend/bd"
)

func DeleteComment2(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	IDComment := vars["IDComment"]
	IDTweet := vars["IDTweet"]

	if len(IDComment) < 1 {
		http.Error(w, "El parámetro ID del comentario es obligatorio", http.StatusBadRequest)
		return
	}

	if len(IDTweet) < 1 {
		http.Error(w, "El parámetro ID del tweet es obligatorio", http.StatusBadRequest)
		return
	}

	err := bd.DeleteComment2(IDComment, IDTweet)
	if err != nil {
		http.Error(w, "Error al intentar borrar el tweet "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)


}