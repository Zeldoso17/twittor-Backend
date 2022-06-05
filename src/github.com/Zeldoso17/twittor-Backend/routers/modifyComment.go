package routers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/Zeldoso17/twittor-Backend/bd"
	"github.com/Zeldoso17/twittor-Backend/models"
)

func ModifyComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	commentID := vars["IDComment"]

	var comment models.Comment
	err := json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		http.Error(w, "Datos Incorrectos"+err.Error(), 400)
		return
	}

	comments := models.CreateComment{
		Mensaje: comment.Mensaje,
	}

	var status bool
	status, err = bd.ModifyComment(comments, commentID, IDUser)

	if err != nil {
		http.Error(w, "Ocurrió un error al intentar modificar el comentario. Reintente nuevamente "+err.Error(), 400)
		return
	}

	if !status {
		http.Error(w, "No se ha logrado modificar el comentario del usuario", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)

}