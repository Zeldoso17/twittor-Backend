package routers

import (
	"encoding/json"
	"net/http"
	"github.com/Zeldoso17/twittor-Backend/bd"
	"github.com/Zeldoso17/twittor-Backend/models"
)

// ReadRelation check if a relation exists between two users
func ReadRelation(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")

	var t models.Relation
	t.UserID = IDUser
	t.UsuarioRelacionID = ID

	var resp models.ResponseReadRelation

	status, err := bd.ReadRelation(t)
	if err != nil || !status {
		resp.Status = false
	}else{
		resp.Status = true
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)

}