package security

import (
	"time"
	
	"github.com/dgrijalva/jwt-go"
	"github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_MID_SPY/models"
)

var jwtKey = []byte("tu_clave_secreta_muy_segura") // Cambia esto en producción!

type Claims struct {
	UserID int    `json:"userId"`
	Email  string `json:"email"`
	RoleID int    `json:"roleId"`
	jwt.StandardClaims
}

func GenerateJWT(user models.Usuarios) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour) // Token expira en 24 horas
	
	claims := &Claims{
		UserID: user.IdUsuarios,
		Email:  user.Email,
		RoleID: user.Rol.Id, // Accede al ID del Rol a través de la relación
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			Issuer:    "SPY-API",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func ValidateToken(tokenString string) (*Claims, error) {
	claims := &Claims{}
	
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	
	if err != nil {
		return nil, err
	}
	
	if !token.Valid {
		return nil, errors.New("token inválido")
	}
	
	return claims, nil
}