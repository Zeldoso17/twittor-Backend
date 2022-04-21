package routers

import (
	"io"
	"net/http"
	"os"

	"github.com/Zeldoso17/twittor-Backend/bd"
)

// ReadAvatar is a funtion that allows to send the avatar to the fontend
func ReadAvatar(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id") // Here we get the id of the user
	if len(ID) < 1 {
		http.Error(w, "El id es obligatorio", http.StatusBadRequest)
		return
	}

	profile, err := bd.ProfileSearch(ID) // Here we search the user
	if err != nil {
		http.Error(w, "Usuario no encontrado! "+err.Error(), http.StatusBadRequest)
		return
	}

	Openfile, err := os.Open("uploads/avatars/" + profile.Avatar) // Here we open the avatar
	if err != nil {
		http.Error(w, "Imagen no encontrada! "+err.Error(), http.StatusBadRequest)
		return
	}

	_, err = io.Copy(w, Openfile) // Here we copy the avatar into the w
	if err != nil {
		http.Error(w, "Error al copiar la imagen! "+err.Error(), http.StatusBadRequest)
	}
}
