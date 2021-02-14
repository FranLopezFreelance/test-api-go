package routes

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/FranLopezFreelance/db"
)

//GetTweets devuelve los tweets de un usuario
func GetTweets(w http.ResponseWriter, r *http.Request){

	userID := r.URL.Query().Get("id")
	pageURL := r.URL.Query().Get("page")

	if len(userID) < 1 {
		http.Error(w, "Debe enviar parámetro id", http.StatusBadRequest)
		return 
	}

	if len(pageURL) < 1 {
		http.Error(w, "Debe enviar parámetro page", http.StatusBadRequest)
		return 
	}

	pageInt , err := strconv.Atoi(pageURL)

	if err != nil {
		http.Error(w, "El parámetro page no es correcto", http.StatusBadRequest)
		return 
	}

	page := int64(pageInt)

	response, status := db.GetTweetsByUser(userID, page)

	if status == false {
		http.Error(w, "Error al leer los tweets", 500)
		return 
	}

	w.Header().Set("COntent-type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(response)

}
