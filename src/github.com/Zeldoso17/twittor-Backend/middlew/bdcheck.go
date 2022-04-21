package middlew

import (
	"net/http"
	"github.com/Zeldoso17/twittor-Backend/bd"
)

// BDcheck is a middleware function that checks if the database is available
func BDcheck(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !bd.ConnectionStatus() {
			http.Error(w, "No DataBase connection", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}