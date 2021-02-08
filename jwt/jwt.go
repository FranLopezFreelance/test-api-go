package jwt

import (
	"time"

	"github.com/FranLopezFreelance/models"
	jwt "github.com/dgrijalva/jwt-go"
)

//Generate genera un token para login
func Generate(user models.User) (string, error) {

	privateKey := [] byte ("@API-test-go-private-key-@FJL")

	payload := jwt.MapClaims{
		"email": user.Email,
		"name": user.Name,
		"lastName": user.LastName,
		"birthdate": user.BirthDate,
		"biography": user.Biography,
		"webSite": user.WebSite,
		"_id": user.ID.Hex(),
		"exp": time.Now().Add(time.Hour * 4).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	tokenStr, err := token.SignedString(privateKey)
	
	if err != nil {
		return tokenStr, err
	}

	return tokenStr, nil
}