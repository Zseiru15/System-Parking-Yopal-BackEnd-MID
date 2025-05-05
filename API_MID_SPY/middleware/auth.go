package middleware

import (
	"net/http"
	"strings"
	"github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_MID_SPY/security"
	"github.com/astaxie/beego/context"
)

func JWTFilter(ctx *context.Context) {
	// Excluir rutas públicas
	publicRoutes := []string{"/v1/auth/login"}
	for _, route := range publicRoutes {
		if ctx.Request.URL.Path == route {
			return
		}
	}

	// Verificar token
	authHeader := ctx.Input.Header("Authorization")
	if authHeader == "" {
		ctx.Output.SetStatus(http.StatusUnauthorized)
		ctx.Output.JSON(map[string]interface{}{
			"success": false,
			"message": "Token de autorización requerido",
		}, false, false)
		return
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	claims, err := security.ValidateToken(tokenString)
	if err != nil {
		ctx.Output.SetStatus(http.StatusUnauthorized)
		ctx.Output.JSON(map[string]interface{}{
			"success": false,
			"message": "Token inválido o expirado",
		}, false, false)
		return
	}

	// Guardar ID de usuario en el contexto
	ctx.Input.SetData("userID", claims.UserId)
}