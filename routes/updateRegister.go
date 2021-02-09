package routes

import (
	"encoding/json"
	"net/http"

	"github.com/FranLopezFreelance/db"
	"github.com/FranLopezFreelance/models"
)

// UpdateProfile modifica los datos del perfil de un usuario
func UpdateProfile(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, "No se pudieron modificar los datos: "+err.Error(), 500)
		return
	}

	var status bool

	status, err = db.UpdateProfile(user, UserID)

	if err != nil {
		http.Error(w, "No se pudieron modificar los datos: "+err.Error(), 500)
		return
	}

	if status == false {
		http.Error(w, "No se pudieron modificar los datos.", 500)
		return
	}

	w.WriteHeader(http.StatusOK)

}
