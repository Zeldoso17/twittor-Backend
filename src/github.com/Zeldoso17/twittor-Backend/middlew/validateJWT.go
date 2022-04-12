package middlew

import (
	"net/http"
	"github.com/Zeldoso17/twittor-Backend/src/github.com/Zeldoso17/twittor-Backend/routers"
)
// ValidateJWT allows to validate the JWT token 
func ValidateJWT(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _, _, err := routers.ProcessToken(r.Header.Get("Authorization"))
		if err != nil {
			http.Error(w, "Error en el Token ! " + err.Error(), http.StatusBadRequest)
			return
		}
		next.ServeHTTP(w, r)
	}
}