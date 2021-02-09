package routes

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/FranLopezFreelance/db"
	"github.com/FranLopezFreelance/jwt"
	"github.com/FranLopezFreelance/models"
)

//Login loguea un usuario
func Login(w http.ResponseWriter, r *http.Request){
	w.Header().Add("Content-Type", "application/json")

	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, "Usuario y/o contraseña no son válidos.", 400)
		return
	}

	if len(user.Email) == 0 {
		http.Error(w, "El email es requerido.", 400)
		return
	}

	document, found := db.TryLogin(user.Email, user.Password)

	if found == false {
		http.Error(w, "Usuario y/o contraseña no son válidos.", 400)
		return
	}

	token, err := jwt.Generate(document)

	if err != nil {
		http.Error(w, "Ha ocurrido un error: "+err.Error(), 500)
		return
	}

	response := models.LoginResponse {
		Token: token,
	}

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)
	
	json.NewEncoder(w).Encode(response)

	expires := time.Now().Add(24 * time.Hour)

	http.SetCookie(w, &http.Cookie{
		Name: "token",
		Value: token,
		Expires: expires,
	})
}
