package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/FranLopezFreelance/middlewares"
	"github.com/FranLopezFreelance/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

//RouterHandlers se encarga del manejo de rutas
func RouterHandlers(){
	router := mux.NewRouter()
	// rutas
	router.HandleFunc("/register", middlewares.CheckDB(routers.Register)).Methods("POST")
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":" + PORT, handler))
}