package routers

import (
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/Zeldoso17/twittor-Backend/src/github.com/Zeldoso17/twittor-Backend/bd"
	"github.com/Zeldoso17/twittor-Backend/src/github.com/Zeldoso17/twittor-Backend/models"
)

// UploadAvatar is a function that allows to upload an banner
func UploadBanner(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("banner")                          // Here we get the avatar from the request
	var extension = strings.Split(handler.Filename, ".")[1]             // Here we get the extension of the banner
	var fileName string = "uploads/banners/" + IDUser + "." + extension // Here we create the name of the banner

	if err != nil {
		http.Error(w, "Error al obtener la imagen del banner"+err.Error(), http.StatusBadRequest)
		return
	}

	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0666) // Here we create the file and open it
	if err != nil {
		http.Error(w, "Error al subir la imagen! "+err.Error(), http.StatusBadRequest)
		return
	}

	_, err = io.Copy(f, file) // Here we copy the banner into the f
	if err != nil {
		http.Error(w, "Error al copiar la imagen! "+err.Error(), http.StatusBadRequest)
		return
	}

	var usuario models.Usuario
	var status bool

	usuario.Banner = IDUser + "." + extension        // Here we set the banner of the user
	status, err = bd.ModifyRegister(usuario, IDUser) // Here we modify the banner from the user
	if err != nil || !status {
		http.Error(w, "Error al intentar guardar el banner! "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json") // Here we set the header of the response
	w.WriteHeader(http.StatusCreated)                  // Here we set the status of the response

}
