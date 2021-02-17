package routes

import (
	"encoding/json"
	"net/http"

	"github.com/FranLopezFreelance/db"
	"github.com/FranLopezFreelance/models"
)

//GetRelation devulve true o false si un usuario sigue a otro
func GetRelation(w http.ResponseWriter, r *http.Request){
	id := r.URL.Query().Get("id")

	var t models.Relation
	t.FollowerID = UserID
	t.FollowedID = id

	var res models.GetRelationResponse
	status, err := db.GetRelation(t)

	if err !=nil || status == false {
		res.Status = false
	} else {
		res.Status = true
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}