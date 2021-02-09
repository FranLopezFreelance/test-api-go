package routes

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/FranLopezFreelance/db"
	"github.com/FranLopezFreelance/models"
)

// NewTweet genera un nuevo tweet
func NewTweet(w http.ResponseWriter, r *http.Request){
	var tweet models.Tweet
	err := json.NewDecoder(r.Body).Decode(&tweet)

	if len(tweet.Message) < 3 {
		http.Error(w, "El mensaje debe contener al menos 3 caractéres.", 400)
	}

	newTweet := models.Tweet{
		UserID: UserID,
		Message: tweet.Message,
		Date: time.Now(),
	}

	_, status, err := db.InsertTweet(newTweet)

	if err != nil {
		http.Error(w, "Error al intentar guardar los datos: "+err.Error(), 500)
	}

	if status == false {
		http.Error(w, "Error al intentar guardar los datos.", 500)
	}

	w.WriteHeader(http.StatusCreated)
}