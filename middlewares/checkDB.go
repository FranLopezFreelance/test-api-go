package middlewares

import (
	"net/http"

	"github.com/FranLopezFreelance/db"
)

// CheckDB chequea que exista la conexion a la base de datos
func CheckDB(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request){
		if db.CheckConnection() == false {
			http.Error(w, "Se ha perdido la conexión con la base de datos.", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
