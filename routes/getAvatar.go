package routes

import (
	"io"
	"net/http"
	"os"

	"github.com/FranLopezFreelance/db"
)

//GetAvatar se encarga de obtener y devolver un avatar
func GetAvatar(w http.ResponseWriter, r *http.Request){
	id := r.URL.Query().Get("id")
	if len(id) < 1 {
		http.Error(w, "Se debe enviar el parámetro id.", http.StatusBadRequest)
		return
	}

	profile, err := db.SearchProfile(id)
	
	if err != nil {
		http.Error(w, "Usuario no encontrado: " + err.Error(), http.StatusBadRequest)
		return
	}

	file, err := os.Open("uploads/avatars/" + profile.Avatar)

	if err != nil {
		http.Error(w, "Imágen no encontrada: " + err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = io.Copy(w, file)

	if err != nil {
		http.Error(w, "Error al copiar la imágen: " + err.Error(), http.StatusInternalServerError)
	}
}