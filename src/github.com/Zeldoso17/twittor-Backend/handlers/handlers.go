package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/Zeldoso17/twittor-Backend/middlew"
	"github.com/Zeldoso17/twittor-Backend/routers"
	"github.com/gofiber/fiber/v2"
	CORS "github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

/* Here I'm setting mi port and i'm listening to my server */
func Managers() {
	router := mux.NewRouter()
	app := fiber.New()
	app.Use(CORS.New())

	router.HandleFunc("/registro", middlew.BDcheck(routers.Register)).Methods("POST")
	router.HandleFunc("/login", middlew.BDcheck(routers.Login)).Methods("POST")
	router.HandleFunc("/verperfil", middlew.BDcheck(middlew.ValidateJWT(routers.ViewProfile))).Methods("GET")
	router.HandleFunc("/modificarPerfil", middlew.BDcheck(middlew.ValidateJWT(routers.ModifyProfile))).Methods("PUT")
	
	router.HandleFunc("/crearTweet", middlew.BDcheck(middlew.ValidateJWT(routers.CreateTweet))).Methods("POST")
	router.HandleFunc("/leerTweets", middlew.BDcheck(middlew.ValidateJWT(routers.ReadTweets))).Methods("GET")
	router.HandleFunc("/eliminarTweet/{IDTweet}", middlew.BDcheck(middlew.ValidateJWT(routers.DeleteTweets))).Methods("DELETE")

	router.HandleFunc("/crearComentario/{IDTweet}", middlew.BDcheck(middlew.ValidateJWT(routers.CreateComment))).Methods("POST")
	router.HandleFunc("/leerComentarios/{IDTweet}", middlew.BDcheck(middlew.ValidateJWT(routers.ReadComments))).Methods("GET")
	router.HandleFunc("/editarComentario/{IDComment}", middlew.BDcheck(middlew.ValidateJWT(routers.ModifyComment))).Methods("PUT")
	router.HandleFunc("/eliminarComentario/{IDComment}", middlew.BDcheck(middlew.ValidateJWT(routers.DeleteComment))).Methods("DELETE")
	router.HandleFunc("/eliminarComentario/{IDComment}/{IDTweet}", middlew.BDcheck(middlew.ValidateJWT(routers.DeleteComment2))).Methods("DELETE")

	router.HandleFunc("/darLike/{IDTweet}", middlew.BDcheck(middlew.ValidateJWT(routers.GiveLike))).Methods("POST")
	router.HandleFunc("/leerLike/{IDTweet}", middlew.BDcheck(middlew.ValidateJWT(routers.ReadLike))).Methods("GET")

	router.HandleFunc("/subirAvatar", middlew.BDcheck(middlew.ValidateJWT(routers.UploadAvatar))).Methods("POST")
	router.HandleFunc("/obtenerAvatar", middlew.BDcheck(routers.ReadAvatar)).Methods("GET")
	router.HandleFunc("/subirBanner", middlew.BDcheck(middlew.ValidateJWT(routers.UploadBanner))).Methods("POST")
	router.HandleFunc("/obtenerBanner", middlew.BDcheck(routers.ReadBanner)).Methods("GET")

	router.HandleFunc("/altaRelacion", middlew.BDcheck(middlew.ValidateJWT(routers.RelationUp))).Methods("POST")
	router.HandleFunc("/bajaRelacion", middlew.BDcheck(middlew.ValidateJWT(routers.DeleteRelation))).Methods("DELETE")
	router.HandleFunc("/consultaRelacion", middlew.BDcheck(middlew.ValidateJWT(routers.ReadRelation))).Methods("GET")

	router.HandleFunc("/listaUsuarios", middlew.BDcheck(middlew.ValidateJWT(routers.ListUsers))).Methods("GET")
	router.HandleFunc("/leoTweetsSeguidores", middlew.BDcheck(middlew.ValidateJWT(routers.ReadFollowersTweets))).Methods("GET")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	//app.Listen(":"+PORT)
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}
