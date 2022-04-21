package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Zeldoso17/twittor-Backend/bd"
)

// ReadTweets is a function that allows to read the tweets of the database
func ReadTweets(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "El parámetro ID es obligatorio", http.StatusBadRequest)
		return
	}

	if len(r.URL.Query().Get("pagina")) < 1 {
		http.Error(w, "El parámetro pagina es obligatorio", http.StatusBadRequest)
		return
	}

	pagina, err := strconv.Atoi(r.URL.Query().Get("pagina"))
	if err != nil {
		http.Error(w, "El parámetro pagina debe ser un entero mayor a 0", http.StatusBadRequest)
		return
	}

	pag := int64(pagina)
	tweets, status := bd.ReadTweets(ID, pag)
	if !status {
		http.Error(w, "Error al leer los tweets", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(tweets)
}
