package routers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Zeldoso17/twittor-Backend/bd"
)

// ListUsers is a function that allows to list all users registered
func ListUsers(w http.ResponseWriter, r *http.Request){
	typeUser := r.URL.Query().Get("type")
	page := r.URL.Query().Get("page")
	search := r.URL.Query().Get("search")

	pagTemp, err := strconv.Atoi(page)
	if err != nil {
		http.Error(w, "Debe enviar el parámetro página como entero mayor a 0", http.StatusBadRequest)
		return
	}

	pag := int64(pagTemp) // Here we convert the string page to an integer

	result, status := bd.ReadAllUsers(IDUser, pag, search, typeUser) // Here we call the function that reads all users registered
	fmt.Println(status)
	if !status {
		http.Error(w, "Error al leer los usuarios", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)

}