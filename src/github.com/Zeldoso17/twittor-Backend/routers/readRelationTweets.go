package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Zeldoso17/twittor-Backend/bd"
)

// ReadFollowersTweet is a function that allows to read the tweets of all my followers
func ReadFollowersTweets(w http.ResponseWriter, r *http.Request) {
	if len(r.URL.Query().Get("pagina")) < 1 {
		http.Error(w, "Debe enviar el parámetro pagina", http.StatusBadRequest)
		return
	}

	pagina, err := strconv.Atoi(r.URL.Query().Get("pagina"))
	if err != nil {
		http.Error(w, "Debe enviar el parámetro pagina como entero mayor a 0", http.StatusBadRequest)
		return
	}

	response, status := bd.ReadFollowersTweet(IDUser, pagina)
	if !status {
		http.Error(w, "Error al leer los tweets", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)

}
