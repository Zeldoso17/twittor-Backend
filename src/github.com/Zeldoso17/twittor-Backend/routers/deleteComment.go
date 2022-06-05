package routers

import (
	"net/http"
	"github.com/gorilla/mux"

	"github.com/Zeldoso17/twittor-Backend/bd"
)

func DeleteComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	IDComment := vars["IDComment"]

	if len(IDComment) < 1 {
		http.Error(w, "El parámetro ID es obligatorio", http.StatusBadRequest)
		return
	}

	err := bd.DeleteComment(IDComment, IDUser)
	if err != nil {
		http.Error(w, "Error al intentar borrar el tweet "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)


}