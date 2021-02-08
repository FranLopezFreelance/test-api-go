package db

import (
	"github.com/FranLopezFreelance/models"
	"golang.org/x/crypto/bcrypt"
)

//TryLogin intenta loguear un usuario
func TryLogin(email string, password string) (models.User, bool) {
	user, found, _ := UserExist(email)
	if found == false {
		return user, false
	}

	passwordBytes := []byte(password)
	passwordDB := []byte(user.Password)
	err := bcrypt.CompareHashAndPassword(passwordDB, passwordBytes)
	if err != nil {
		return user, false
	}

	return user, true
}
