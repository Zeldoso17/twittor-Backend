package routers

import (
	"net/http"
	"github.com/Zeldoso17/twittor-Backend/bd"
)

// DeleteTweet is a function that allows to delete a tweet
func DeleteTweets(w http.ResponseWriter, r *http.Request){
	ID := r.URL.Query().Get("id") // Here we get the id of the tweet from the url
	if len(ID) < 1 {
		http.Error(w, "El parámetro ID es obligatorio", http.StatusBadRequest)
		return
	}

	err := bd.DeleteTweet(ID, IDUser) // Here we delete the tweet
	if err != nil {
		http.Error(w, "Error al intentar borrar el tweet "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json") // Here we set the header of the response
	w.WriteHeader(http.StatusCreated) // Here we set the status of the response
}