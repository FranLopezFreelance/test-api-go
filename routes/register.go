package routes

import (
	"encoding/json"
	"net/http"

	"github.com/FranLopezFreelance/db"
	"github.com/FranLopezFreelance/models"
)

//Register registra un usuario
func Register(w http.ResponseWriter, r *http.Request){
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	
	if err != nil {
		http.Error(w, "error en los datos recibidos: " + err.Error(), 400)
		return
	}

	if len(user.Email) == 0 {
		http.Error(w, "el email es requerido.", 400)
		return
	}

	if len(user.Password) < 6 {
		http.Error(w, "la contraseña debe ser de al menos 6 caractéres.", 400)
		return
	}

	_, found, _ := db.UserExist(user.Email)

	if found == true {
		http.Error(w, "el usuario ya existe", 400)
		return
	}

	_, status, err := db.Register(user)

	if err != nil {
		http.Error(w, "no se pudo realizar el registro: " + err.Error(), 500)
		return
	}

	if status == false {
		http.Error(w, "no se pudo realizar el registro", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
