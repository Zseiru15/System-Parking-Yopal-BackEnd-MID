package middleware

import (
	"strings"
	
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web/context"
	"github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_MID_SPY/security"
)

// Auth middleware
var Auth = func(ctx *context.Context) {
	authHeader := ctx.Input.Header("Authorization")

	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		logs.Warn("Token no proporcionado o mal formado")
		ctx.Output.SetStatus(401)
		ctx.Output.JSON(map[string]interface{}{"Success": false, "Message": "Token no proporcionado"}, false, true)
		return
	}

	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	claims, err := security.ValidateToken(tokenString)
	if err != nil || claims.UserID == 0 {
		logs.Warn("Token inválido")
		ctx.Output.SetStatus(401)
		ctx.Output.JSON(map[string]interface{}{"Success": false, "Message": "Token inválido"}, false, true)
		return
	}

	// El middleware termina aquí sin necesidad de ctx.Next()
}
