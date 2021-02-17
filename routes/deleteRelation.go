package routes

import (
	"net/http"

	"github.com/FranLopezFreelance/db"
	"github.com/FranLopezFreelance/models"
)

//DeleteRelation se encarga de crear una relación entre usuarios
func DeleteRelation(w http.ResponseWriter, r *http.Request){
	id := r.URL.Query().Get("id")
	if len(id) < 1 {
		http.Error(w, "El parámetro id es necesario.", http.StatusBadRequest)
		return
	}

	var t models.Relation
	t.FollowerID = UserID
	t.FollowedID = id

	status, err := db.DeleteRelation(t)

	if err != nil {
		http.Error(w, "Ocurrió un error al intentar actualizar.", http.StatusBadRequest)
		return
	}
	
	if status == false {
		http.Error(w, "Ocurrió un error al intentar actualizar.", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)

}