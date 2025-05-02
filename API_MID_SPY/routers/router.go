// @APIVersion 1.0.0
// @Title SPY API
// @Description API para el Sistema Parking Yopal
// @Contact sneyder@example.com
// @TermsOfServiceUrl http://spy.example.com/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/astaxie/beego"
	"github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_MID_SPY/controllers"
	"github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_MID_SPY/middleware"
)

func init() {
	// Aplicar middleware de autenticación a todas las rutas
	beego.InsertFilter("*", beego.BeforeRouter, middleware.AuthMiddleware)

	ns := beego.NewNamespace("/v1",
		// Rutas públicas de autenticación
		beego.NSNamespace("/auth",
			beego.NSRouter("/login", &controllers.AuthController{}, "post:Login"),
			beego.NSRouter("/register", &controllers.AuthController{}, "post:Register"),
		),

		// Rutas protegidas
		beego.NSNamespace("/comentarios",
			beego.NSInclude(
				&controllers.ComentariosController{},
			),
		),
		beego.NSNamespace("/pagos",
			beego.NSInclude(
				&controllers.PagosController{},
			),
		),
		beego.NSNamespace("/parqueaderos",
			beego.NSInclude(
				&controllers.ParqueaderosController{},
			),
		),
		beego.NSNamespace("/usuarios",
			beego.NSInclude(
				&controllers.UsuariosController{},
			),
		),
		beego.NSNamespace("/vehiculos",
			beego.NSInclude(
				&controllers.VehiculosController{},
			),
		),
	)

	beego.AddNamespace(ns)
}