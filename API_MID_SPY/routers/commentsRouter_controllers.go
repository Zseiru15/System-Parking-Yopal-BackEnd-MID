package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_MID_SPY/controllers:AuthController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_MID_SPY/controllers:AuthController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_MID_SPY/controllers:AuthController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_MID_SPY/controllers:AuthController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_MID_SPY/controllers:AuthController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_MID_SPY/controllers:AuthController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_MID_SPY/controllers:AuthController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_MID_SPY/controllers:AuthController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_MID_SPY/controllers:AuthController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_MID_SPY/controllers:AuthController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_MID_SPY/controllers:ComentariosController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_MID_SPY/controllers:ComentariosController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_MID_SPY/controllers:ComentariosController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_MID_SPY/controllers:ComentariosController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_MID_SPY/controllers:ComentariosController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_MID_SPY/controllers:ComentariosController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_MID_SPY/controllers:ComentariosController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_MID_SPY/controllers:ComentariosController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_MID_SPY/controllers:ComentariosController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_MID_SPY/controllers:ComentariosController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_MID_SPY/controllers:PagosController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_MID_SPY/controllers:PagosController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_MID_SPY/controllers:PagosController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_MID_SPY/controllers:PagosController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_MID_SPY/controllers:PagosController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_MID_SPY/controllers:PagosController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_MID_SPY/controllers:PagosController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_MID_SPY/controllers:PagosController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_MID_SPY/controllers:PagosController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_MID_SPY/controllers:PagosController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_MID_SPY/controllers:ParqueaderosController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_MID_SPY/controllers:ParqueaderosController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_MID_SPY/controllers:ParqueaderosController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_MID_SPY/controllers:ParqueaderosController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_MID_SPY/controllers:ParqueaderosController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_MID_SPY/controllers:ParqueaderosController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_MID_SPY/controllers:ParqueaderosController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_MID_SPY/controllers:ParqueaderosController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_MID_SPY/controllers:ParqueaderosController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_MID_SPY/controllers:ParqueaderosController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_MID_SPY/controllers:ParqueaderosController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_MID_SPY/controllers:ParqueaderosController"],
        beego.ControllerComments{
            Method: "GetParqueaderoDelEmpleado",
            Router: "/parqueadero-empleado/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_MID_SPY/controllers:ParqueaderosController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_MID_SPY/controllers:ParqueaderosController"],
        beego.ControllerComments{
            Method: "GetParqueaderosPorEmpleado",
            Router: "/parqueaderos-empleado/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_MID_SPY/controllers:ParqueaderosController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_MID_SPY/controllers:ParqueaderosController"],
        beego.ControllerComments{
            Method: "PostPromocion",
            Router: "/promociones",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_MID_SPY/controllers:ParqueaderosController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_MID_SPY/controllers:ParqueaderosController"],
        beego.ControllerComments{
            Method: "GetPromocionesPorParqueadero",
            Router: "/promociones/:idParqueadero",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_MID_SPY/controllers:ParqueaderosController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_MID_SPY/controllers:ParqueaderosController"],
        beego.ControllerComments{
            Method: "GetByUsuario",
            Router: "/usuario/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_MID_SPY/controllers:UsuariosController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_MID_SPY/controllers:UsuariosController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_MID_SPY/controllers:UsuariosController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_MID_SPY/controllers:UsuariosController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_MID_SPY/controllers:UsuariosController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_MID_SPY/controllers:UsuariosController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_MID_SPY/controllers:UsuariosController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_MID_SPY/controllers:UsuariosController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_MID_SPY/controllers:UsuariosController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_MID_SPY/controllers:UsuariosController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_MID_SPY/controllers:UsuariosController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_MID_SPY/controllers:UsuariosController"],
        beego.ControllerComments{
            Method: "GetUsuarioPorIdentificacion",
            Router: "/identificacion/:identificacion",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_MID_SPY/controllers:UsuariosController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_MID_SPY/controllers:UsuariosController"],
        beego.ControllerComments{
            Method: "GetByTrabajadores",
            Router: "/usuario/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_MID_SPY/controllers:VehiculosController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_MID_SPY/controllers:VehiculosController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_MID_SPY/controllers:VehiculosController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_MID_SPY/controllers:VehiculosController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_MID_SPY/controllers:VehiculosController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_MID_SPY/controllers:VehiculosController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_MID_SPY/controllers:VehiculosController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_MID_SPY/controllers:VehiculosController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_MID_SPY/controllers:VehiculosController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_MID_SPY/controllers:VehiculosController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_MID_SPY/controllers:VehiculosController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_MID_SPY/controllers:VehiculosController"],
        beego.ControllerComments{
            Method: "GetByUsuario",
            Router: "/usuario/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
