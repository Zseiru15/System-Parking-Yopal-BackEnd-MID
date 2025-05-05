package routers

import (
	"github.com/astaxie/beego"
	"github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_MID_SPY/controllers"
	"github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_MID_SPY/middleware"
)

func init() {

	// Aplicar middleware JWT a todas las rutas
	beego.InsertFilter("*", beego.BeforeRouter, middleware.JWTFilter)

	// Resto de tus rutas protegidas...
	ns := beego.NewNamespace("/v1",
		// Rutas de autenticación (públicas)
		beego.NSRouter("/auth/login", &controllers.AuthController{}, "post:Login"),
		beego.NSRouter("/auth/register", &controllers.AuthController{}, "post:Register"),
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
