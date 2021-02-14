package routes

import (
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/FranLopezFreelance/db"
	"github.com/FranLopezFreelance/models"
)

//UploadBanner se encarga de subir un banner al servidor y guardar la referencia en la base de datos
func UploadBanner(w http.ResponseWriter, r *http.Request){
	file, handler, err := r.FormFile("banner")
	var extension = strings.Split(handler.Filename, ".")[1]
	var location = "uploads/banners/"
	var fileName string = location + UserID + "." + extension
	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, "Error al subir la imágen: " + err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = io.Copy(f, file)
	if err != nil {
		http.Error(w, "Error al copiar la imágen: " + err.Error(), http.StatusInternalServerError)
		return
	}
	var user models.User
	var status bool
	user.Banner = UserID + "." + extension
	status, err = db.UpdateProfile(user, UserID)

	if err != nil || status == false {
		http.Error(w, "Error al grabar en la DB: " + err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

