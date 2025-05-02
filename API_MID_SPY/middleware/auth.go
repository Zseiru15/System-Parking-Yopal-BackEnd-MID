// middleware/auth.go
package middleware

import (
	"net/http"
	"strings"

	"github.com/astaxie/beego/context"
	"github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_MID_SPY/security"
)

func AuthMiddleware(ctx *context.Context) {
	// Lista de rutas públicas que no requieren autenticación
	publicRoutes := []string{
		"/v1/auth/login",
		"/v1/auth/register",
		"/v1/auth", // Ruta base por si acaso
	}

	// Verificar si la ruta actual es pública
	currentPath := ctx.Input.URL()
	for _, route := range publicRoutes {
		if strings.HasPrefix(currentPath, route) {
			return // No aplicar autenticación a rutas públicas
		}
	}

	// Resto de la lógica de autenticación...
	authHeader := ctx.Input.Header("Authorization")
	if authHeader == "" {
		ctx.Output.SetStatus(http.StatusUnauthorized)
		ctx.Output.JSON(map[string]string{"error": "Token de autorización requerido"}, false, false)
		return
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	if tokenString == authHeader {
		ctx.Output.SetStatus(http.StatusUnauthorized)
		ctx.Output.JSON(map[string]string{"error": "Formato de token inválido. Use 'Bearer <token>'"}, false, false)
		return
	}

	_, err := security.ValidateToken(tokenString)
	if err != nil {
		ctx.Output.SetStatus(http.StatusUnauthorized)
		ctx.Output.JSON(map[string]string{"error": "Token inválido o expirado"}, false, false)
		return
	}
}

func AdminMiddleware(ctx *context.Context) {
    claims, err := security.ValidateToken(getTokenFromHeader(ctx))
    if err != nil || claims.RoleID != 1 { // 1 = ID de admin
        ctx.Output.SetStatus(http.StatusForbidden)
        ctx.Output.JSON(map[string]string{"error": "Acceso no autorizado"}, false, false)
        return
    }
}