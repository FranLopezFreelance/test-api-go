package middlewares

import (
	"net/http"

	"github.com/FranLopezFreelance/routes"
)

// JWTValidate valida el token de los request
func JWTValidate(next http.HandlerFunc) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		_, _, _, err := routes.CheckToken(r.Header.Get("Authorization"))

		if err != nil {
			http.Error(w, "Token inválido: "+err.Error(), 401)
			return
		}

		next.ServeHTTP(w, r)
	}
}
