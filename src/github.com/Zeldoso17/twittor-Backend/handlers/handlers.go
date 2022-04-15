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
	router.HandleFunc("/crearTweet", middlew.BDcheck(middlew.ValidateJWT(routers.CreateTweet))).Methods("POST")
	router.HandleFunc("/leerTweets", middlew.BDcheck(middlew.ValidateJWT(routers.ReadTweets))).Methods("GET")
	router.HandleFunc("/eliminarTweet", middlew.BDcheck(middlew.ValidateJWT(routers.DeleteTweets))).Methods("DELETE")

	router.HandleFunc("/subirAvatar", middlew.BDcheck(middlew.ValidateJWT(routers.UploadAvatar))).Methods("POST")
	router.HandleFunc("/obtenerAvatar", middlew.BDcheck(routers.ReadAvatar)).Methods("GET")
	router.HandleFunc("/subirBanner", middlew.BDcheck(middlew.ValidateJWT(routers.UploadBanner))).Methods("POST")
	router.HandleFunc("/obtenerBanner", middlew.BDcheck(routers.ReadBanner)).Methods("GET")

	router.HandleFunc("/altaRelacion", middlew.BDcheck(middlew.ValidateJWT(routers.RelationUp))).Methods("POST")
	router.HandleFunc("/bajaRelacion", middlew.BDcheck(middlew.ValidateJWT(routers.DeleteRelation))).Methods("DELETE")
	router.HandleFunc("/consultaRelacion", middlew.BDcheck(middlew.ValidateJWT(routers.ReadRelation))).Methods("GET")

	router.HandleFunc("/listaUsuarios", middlew.BDcheck(middlew.ValidateJWT(routers.ListUsers))).Methods("GET")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}

