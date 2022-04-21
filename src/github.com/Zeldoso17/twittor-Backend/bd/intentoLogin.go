package bd

import (
	"github.com/Zeldoso17/twittor-Backend/models"
	"golang.org/x/crypto/bcrypt"
)

// LoginTry is a function that checks if the user is registered and if the password is correct
func LoginTry(email string, password string) (models.Usuario, bool) {
	user, found, _ := UserAlreadyExist(email)
	if !found {
		return user, found
	}

	passwordBytes := []byte(password)
	passwordBD := []byte(user.Password) // Here we are getting the password from the database

	err := bcrypt.CompareHashAndPassword(passwordBD, passwordBytes) // Here we are comparing the password from the database with the password that the user entered
	if err != nil {
		return user, false
	}
	return user, true
}
