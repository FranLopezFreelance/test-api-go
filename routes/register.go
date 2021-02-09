package routes

import (
	"encoding/json"
	"net/http"

	"github.com/FranLopezFreelance/db"
	"github.com/FranLopezFreelance/models"
)

// Register registra un usuario
func Register(w http.ResponseWriter, r *http.Request){
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	
	if err != nil {
		http.Error(w, "Error en los datos recibidos: " + err.Error(), 400)
		return
	}

	if len(user.Email) == 0 {
		http.Error(w, "El email es requerido.", 400)
		return
	}

	if len(user.Password) < 6 {
		http.Error(w, "La contraseña debe ser de al menos 6 caractéres.", 400)
		return
	}

	_, found, _ := db.UserExist(user.Email)

	if found == true {
		http.Error(w, "El usuario ya existe.", 400)
		return
	}

	_, status, err := db.Register(user)

	if err != nil {
		http.Error(w, "nN se pudo realizar el registro: " + err.Error(), 500)
		return
	}

	if status == false {
		http.Error(w, "No se pudo realizar el registro", 500)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
