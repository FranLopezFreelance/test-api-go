package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/FranLopezFreelance/middlewares"
	"github.com/FranLopezFreelance/routes"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// RouterHandlers se encarga del manejo de rutas
func RouterHandlers(){

	router := mux.NewRouter()

	// API ROUTES
	router.HandleFunc("/api/auth/register", 
		middlewares.CheckDB(routes.Register),
	).Methods("POST")

	router.HandleFunc("/api/auth/login", 
		middlewares.CheckDB(routes.Login),
	).Methods("POST")

	router.HandleFunc("/api/viewProfile", 
		middlewares.CheckDB(
			middlewares.JWTValidate(routes.ViewProfile),
		),
	).Methods("GET")

	router.HandleFunc("/api/updateProfile", 
		middlewares.CheckDB(
			middlewares.JWTValidate(routes.UpdateProfile),
		),
	).Methods("PUT")

	router.HandleFunc("/api/newTweet", 
		middlewares.CheckDB(
			middlewares.JWTValidate(routes.NewTweet),
		),
	).Methods("POST")
	
	// LISTEN AND SERVE
	PORT := os.Getenv("PORT")
	
	if PORT == "" {
		PORT = "8080"
	}
	
	handler := cors.AllowAll().Handler(router)
	
	log.Fatal(http.ListenAndServe(":" + PORT, handler))
}
