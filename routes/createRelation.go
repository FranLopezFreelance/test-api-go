package routes

import (
	"net/http"

	"github.com/FranLopezFreelance/db"
	"github.com/FranLopezFreelance/models"
)

//CreateRelation se encarga de crear una relación entre usuarios
func CreateRelation(w http.ResponseWriter, r *http.Request){
	id := r.URL.Query().Get("id");
	if len(id)<1 {
		http.Error(w, "El parámetro id es necesario.", http.StatusBadRequest)
		return
	}

	var t models.Relation
	t.FollowerID = UserID
	t.FollowedID = id

	status, err := db.CreateRelation(t)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar seguir.", http.StatusBadRequest)
		return
	}
	
	if status == false {
		http.Error(w, "Ocurrió un error al intentar seguir.", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}	
