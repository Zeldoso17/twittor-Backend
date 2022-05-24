package routers

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"

	"github.com/Zeldoso17/twittor-Backend/bd"
)

func ReadLike(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	IDTweet := vars["IDTweet"]

	status, _ := bd.ReadLike(IDUser, IDTweet)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(status)
}