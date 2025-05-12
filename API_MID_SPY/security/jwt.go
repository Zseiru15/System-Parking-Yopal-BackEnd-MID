package security

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var jwtKey = []byte("mi_clave_secreta") // usa una más segura en producción

func GenerarToken(email string) (string, error) {
	claims := jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(jwtKey)
}
