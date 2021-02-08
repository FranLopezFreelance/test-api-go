package routers

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
		http.Error(w, "Ya existe ese Usuario.", 400)
		return
	}

	_, status, err := db.Register(user)
	if err != nil {
		http.Error(w, "Ha ocurrido un error en el registro: " + err.Error(), 400)
		return
	}
	if status == false {
		http.Error(w, "No se pudieron guardar los datos.", 400)
		return
	}

	return
}