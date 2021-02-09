package routes

import (
	"errors"
	"strings"

	"github.com/FranLopezFreelance/db"
	"github.com/FranLopezFreelance/models"
	"github.com/dgrijalva/jwt-go"
)

// UserID es el id del usuario logueado
var UserID string

// Email es el email del usuario logueado
var Email string

// CheckToken se encarga de validar el token
func CheckToken(bearerToken string) (*models.Claims, bool, string, error) {
	privateKey := [] byte ("@API-test-go-private-key-@FJL")
	claims := &models.Claims{}
	tokenSplit := strings.Split(bearerToken, "Bearer")
	if len(tokenSplit) != 2 {
		return claims, false, string(""), errors.New("token inválido")
	} 

	token := strings.TrimSpace(tokenSplit[1])

	tokenClaims, err := jwt.ParseWithClaims(token, claims, func (tk *jwt.Token) (interface {}, error){
		return privateKey, nil
	})

	if err == nil {
		_, found, _ := db.UserExist(claims.Email)
		if found == true {
			UserID = claims.ID.Hex()
			Email = claims.Email
		}

		return claims, found, UserID, nil
	}

	if !tokenClaims.Valid {
		return claims, false, string(""), errors.New("token inválido")
	}

	return claims, false, string(""), err
}
