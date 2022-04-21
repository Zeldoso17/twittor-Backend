package routers

import (
	"io"
	"net/http"
	"os"

	"github.com/Zeldoso17/twittor-Backend/bd"
)

// ReadBanner is a funtion that allows to send the banner to the fontend
func ReadBanner(w http.ResponseWriter, r *http.Request) {
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

	Openfile, err := os.Open("uploads/banners/" + profile.Banner) // Here we open the banner
	if err != nil {
		http.Error(w, "Imagen no encontrada! "+err.Error(), http.StatusBadRequest)
		return
	}

	_, err = io.Copy(w, Openfile) // Here we copy the banner into the w
	if err != nil {
		http.Error(w, "Error al copiar la imagen! "+err.Error(), http.StatusBadRequest)
	}
}
