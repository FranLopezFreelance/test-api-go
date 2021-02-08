package routes

import (
	"encoding/json"
	"net/http"

	"github.com/FranLopezFreelance/db"
)

//ViewProfile devuelve los datos de perfil de usuario
func ViewProfile(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if len(id) == 0 {
		http.Error(w, "parámetro id es requerido", 400)
		return
	}

	profile, err := db.SearchProfile(id)

	if err != nil {
		http.Error(w, "no se pudo acceder a los datos de perfil: "+err.Error(), 500)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(profile)
}