package main

import (
	"log"
	"github.com/Zeldoso17/twittor-Backend/bd"
	"github.com/Zeldoso17/twittor-Backend/handlers"
)

func main() {
	if !bd.ConnectionStatus() {
		log.Fatal("No Database Connection")
		return
	}
	handlers.Managers()
}