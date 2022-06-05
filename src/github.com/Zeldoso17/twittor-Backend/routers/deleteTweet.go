package routers

import (
	"net/http"
	"github.com/gorilla/mux"

	"github.com/Zeldoso17/twittor-Backend/bd"
)

// DeleteTweet is a function that allows to delete a tweet
func DeleteTweets(w http.ResponseWriter, r *http.Request){
	vars := mux .Vars(r) 

	IDTweet := vars["IDTweet"]

	if len(IDTweet) < 1 {
		http.Error(w, "El parámetro ID es obligatorio", http.StatusBadRequest)
		return
	}

	err := bd.DeleteTweet(IDTweet, IDUser) // Here we delete the tweet
	if err != nil {
		http.Error(w, "Error al intentar borrar el tweet "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json") // Here we set the header of the response
	w.WriteHeader(http.StatusCreated) // Here we set the status of the response
}