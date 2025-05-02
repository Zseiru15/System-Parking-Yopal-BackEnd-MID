// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/astaxie/beego"
	"github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_MID_SPY/controllers"
)

func init() {
	ns := beego.NewNamespace("/v1",
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
	beego.NSNamespace("/auth",
		beego.NSRouter("/login", &controllers.AuthController{}, "post:Login"),
		beego.NSRouter("/register", &controllers.AuthController{}, "post:Register"),
	),
)

beego.AddNamespace(ns)
}
