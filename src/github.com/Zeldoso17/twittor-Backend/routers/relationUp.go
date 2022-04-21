package routers

import (
	"net/http"
	"github.com/Zeldoso17/twittor-Backend/bd"
	"github.com/Zeldoso17/twittor-Backend/models"
)

// RelationUp is a function that allows to make a relation between two users
func RelationUp(w http.ResponseWriter, r *http.Request){
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "El id es requerido", http.StatusBadRequest)
		return
	}

	var t models.Relation
	t.UserID = IDUser
	t.UsuarioRelacionID = ID

	status, err := bd.InsertRelation(t)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar realizar el registro de la relación "+err.Error(), http.StatusBadRequest)
		return
	}

	if !status {
		http.Error(w, "No se ha logrado insertar la relación", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)

}