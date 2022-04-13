package handlers

import (
	"log"
	"net/http"
	"os"
	"github.com/Zeldoso17/twittor-Backend/src/github.com/Zeldoso17/twittor-Backend/middlew"
	"github.com/Zeldoso17/twittor-Backend/src/github.com/Zeldoso17/twittor-Backend/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

/* Here I'm setting mi port and i'm listening to my server */
func Managers(){
	router := mux.NewRouter()

	router.HandleFunc("/registro", middlew.BDcheck(routers.Register)).Methods("POST")
	router.HandleFunc("/login", middlew.BDcheck(routers.Login)).Methods("POST")
	router.HandleFunc("/verperfil", middlew.BDcheck(middlew.ValidateJWT(routers.ViewProfile))).Methods("GET")
	router.HandleFunc("/modificarPerfil", middlew.BDcheck(middlew.ValidateJWT(routers.ModifyProfile))).Methods("PUT")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}

