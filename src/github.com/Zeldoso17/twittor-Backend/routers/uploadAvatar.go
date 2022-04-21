package routers

import (
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/Zeldoso17/twittor-Backend/bd"
	"github.com/Zeldoso17/twittor-Backend/models"
)

// UploadAvatar is a function that allows to upload an avatar
func UploadAvatar(w http.ResponseWriter, r *http.Request){
	file, handler, err := r.FormFile("avatar") // Here we get the avatar from the request
	var extension = strings.Split(handler.Filename, ".")[1] // Here we get the extension of the avatar
	var fileName string = "uploads/avatars/" + IDUser + "." + extension // Here we create the name of the avatar

	if err != nil {
		http.Error(w, "Error al obtener la imagen del avatar"+err.Error(), http.StatusBadRequest)
		return
	}

	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0666) // Here we create the file and open it
	if err != nil {
		http.Error(w, "Error al subir la imagen! "+err.Error(), http.StatusBadRequest)
		return
	}

	_, err = io.Copy(f, file) // Here we copy the avatar into the f
	if err != nil {
		http.Error(w, "Error al copiar la imagen! "+err.Error(), http.StatusBadRequest)
		return
	}

	var usuario models.Usuario
	var status bool

	usuario.Avatar = IDUser + "." + extension // Here we set the avatar of the user
	status, err = bd.ModifyRegister(usuario, IDUser) // Here we modify the Avatar from the user
	if err != nil || !status {
		http.Error(w, "Error al intentar guardar el avatar! "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json") // Here we set the header of the response
	w.WriteHeader(http.StatusCreated) // Here we set the status of the response

}
