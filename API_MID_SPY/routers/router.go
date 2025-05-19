package routers

import (
	"github.com/astaxie/beego"
	"github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_MID_SPY/controllers"
)

func init() {

	// Resto de tus rutas protegidas...
	ns := beego.NewNamespace("/v1",
		// Rutas de autenticación (públicas)
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
			beego.NSRouter("/usuario/:id", &controllers.ParqueaderosController{}, "get:GetByUsuario"),
			beego.NSInclude(
				&controllers.ParqueaderosController{}),
		),

		beego.NSNamespace("/usuarios",
			beego.NSRouter("/login", &controllers.UsuariosController{}, "post:Login"),
			beego.NSInclude(
				&controllers.UsuariosController{},
			),
		),

		beego.NSNamespace("/vehiculos",
			beego.NSRouter("/usuario/:id", &controllers.VehiculosController{}, "get:GetByUsuario"),
			beego.NSInclude(
				&controllers.VehiculosController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
