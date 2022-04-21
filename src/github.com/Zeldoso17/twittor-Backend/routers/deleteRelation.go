package routers

import (
	"net/http"
	"github.com/Zeldoso17/twittor-Backend/bd"
	"github.com/Zeldoso17/twittor-Backend/models"
)

// DeleteRelation is a function that deletes a relation
func DeleteRelation(w http.ResponseWriter, r *http.Request){
	ID := r.URL.Query().Get("id")
	var t models.Relation
	t.UserID = IDUser
	t.UsuarioRelacionID = ID

	status, err := bd.DeleteRelation(t)
	if !status {
		http.Error(w, "Error al eliminar la relación", http.StatusBadRequest)
		return
	}
	if err != nil {
		http.Error(w, "No se ha logrado borrar la relación "+err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}