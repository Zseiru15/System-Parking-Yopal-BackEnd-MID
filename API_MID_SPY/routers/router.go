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
		beego.NSRouter("/actividades-usuario/:id", &controllers.ActividadController{}, "get:GetActividadesPorUsuario"),
		
		beego.NSNamespace("/comentarios",
			beego.NSInclude(
				&controllers.ComentariosController{},
			),
		),
		beego.NSNamespace("/pagos",
			beego.NSRouter("/parqueadero/:id", &controllers.PagosController{}, "get:GetByEstacionamientoId"),
			beego.NSRouter("/usuario/:id", &controllers.PagosController{}, "get:GetByUsuarioId"),
			beego.NSInclude(
				&controllers.PagosController{},
			),
		),
		beego.NSNamespace("/parqueaderos",
			beego.NSRouter("/usuario/:id", &controllers.ParqueaderosController{}, "get:GetByUsuario"),
			beego.NSRouter("/empleo/:id", &controllers.ParqueaderosController{}, "get:GetParqueaderoDelEmpleado"),
			beego.NSRouter("/promociones/:idParqueadero", &controllers.ParqueaderosController{}, "get:GetPromocionesPorParqueadero"),
			beego.NSRouter("/registrarPromocion/:idParqueadero", &controllers.ParqueaderosController{}, "post:PostPromocion"),
			beego.NSRouter("/desactivar/:id", &controllers.ParqueaderosController{}, "put:DesactivarParqueadero"),
			beego.NSInclude(
				&controllers.ParqueaderosController{}),
		),
		beego.NSNamespace("/usuarios",
			beego.NSRouter("/login", &controllers.UsuariosController{}, "post:Login"),
			beego.NSRouter("/trabajador/:id", &controllers.UsuariosController{}, "get:GetByTrabajadores"),
			beego.NSRouter("/identificacion/:identificacion", &controllers.UsuariosController{}, "get:GetUsuarioPorIdentificacion"),
			beego.NSRouter("/activar-membresia", &controllers.UsuariosController{}, "post:ActivarMembresia"),
			beego.NSRouter("/desactivar-membresias", &controllers.UsuariosController{}, "get:DesactivarMembresiasVencidas"),
			beego.NSInclude(
				&controllers.UsuariosController{},
			),
		),
		beego.NSNamespace("/vehiculos",
			beego.NSRouter("/usuario/:id", &controllers.VehiculosController{}, "get:GetByUsuario"),
			beego.NSRouter("/desactivar/:id", &controllers.VehiculosController{}, "put:DesactivarVehiculo"),
			beego.NSInclude(
				&controllers.VehiculosController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
