package security

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/astaxie/beego/context"
	"github.com/golang-jwt/jwt/v4"
)

// jwtSecret se carga desde la variable de entorno JWT_SECRET
var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

// CustomClaims define los claims personalizados que queremos incluir en el token.
// Aquí se incluyen el UserID, Email y Role. Puedes agregar más campos según lo necesites.
type CustomClaims struct {
	Idusuarios int    `json:"user_id"`
	Email  string `json:"email"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

// GenerateJWT genera un token JWT para un usuario dado.
// El token expira en 24 horas
func GenerateJWT(Id_usuarios int, Email, IdRolesFk string) (string, error) {
	claims := CustomClaims{
		Idusuarios: Id_usuarios,
		Email:  Email,
		Role:   IdRolesFk,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "SPY",
		},
	}

	// Crear el token usando el método de firma HS256
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Firmar el token con la clave secreta
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", fmt.Errorf("error generando token: %v", err)
	}
	return tokenString, nil
}

// verifica si se envía un token JWT válido.
func jwtFilter(ctx *context.Context) {
	// Se espera que el token se envíe en el header "Authorization"
	tokenStr := ctx.Input.Header("Authorization")
	if tokenStr == "" {
		ctx.Output.SetStatus(http.StatusUnauthorized)
		ctx.Output.Body([]byte("No se proporcionó token"))
		return
	}

	// Validar el token utilizando la función ValidarJWT
	_, err := ValidarJWT(tokenStr)
	if err != nil {
		ctx.Output.SetStatus(http.StatusUnauthorized)
		ctx.Output.Body([]byte("Token inválido o expirado"))
		return
	}
	// Si el token es válido, el filtro permite continuar con la solicitud.
}

// valida un token JWT y retorna los claims si es válido.
// Se asegura que el método de firma sea el esperado.
func ValidarJWT(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Verificar el método de firma
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("método de firma inesperado: %v", token.Header["alg"])
		}
		return jwtSecret, nil
	})
	if err != nil {
		return nil, err
	}

	// Si el token es válido, se retornan los claims
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, fmt.Errorf("token inválido")
}