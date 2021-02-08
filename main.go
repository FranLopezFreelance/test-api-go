package main

import (
	"log"

	"github.com/FranLopezFreelance/db"
	"github.com/FranLopezFreelance/handlers"
)

func main() {
	if db.CheckConnection() == false {
		log.Fatal("No hay conexión a la base de datos")
	}
	handlers.RouterHandlers()
}