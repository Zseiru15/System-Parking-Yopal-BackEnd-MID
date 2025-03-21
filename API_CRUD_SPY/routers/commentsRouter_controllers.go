package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_CRUD_SPY/controllers:AdministradoresEmpleadosController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_CRUD_SPY/controllers:AdministradoresEmpleadosController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_CRUD_SPY/controllers:AdministradoresEmpleadosController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_CRUD_SPY/controllers:AdministradoresEmpleadosController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_CRUD_SPY/controllers:AdministradoresEmpleadosController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_CRUD_SPY/controllers:AdministradoresEmpleadosController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_CRUD_SPY/controllers:AdministradoresEmpleadosController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_CRUD_SPY/controllers:AdministradoresEmpleadosController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_CRUD_SPY/controllers:AdministradoresEmpleadosController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_CRUD_SPY/controllers:AdministradoresEmpleadosController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_CRUD_SPY/controllers:AdministradoresEstacionamientosController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_CRUD_SPY/controllers:AdministradoresEstacionamientosController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_CRUD_SPY/controllers:AdministradoresEstacionamientosController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_CRUD_SPY/controllers:AdministradoresEstacionamientosController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_CRUD_SPY/controllers:AdministradoresEstacionamientosController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_CRUD_SPY/controllers:AdministradoresEstacionamientosController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_CRUD_SPY/controllers:AdministradoresEstacionamientosController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_CRUD_SPY/controllers:AdministradoresEstacionamientosController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_CRUD_SPY/controllers:AdministradoresEstacionamientosController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_CRUD_SPY/controllers:AdministradoresEstacionamientosController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_CRUD_SPY/controllers:CargosController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_CRUD_SPY/controllers:CargosController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_CRUD_SPY/controllers:CargosController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_CRUD_SPY/controllers:CargosController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_CRUD_SPY/controllers:CargosController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_CRUD_SPY/controllers:CargosController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_CRUD_SPY/controllers:CargosController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_CRUD_SPY/controllers:CargosController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_CRUD_SPY/controllers:CargosController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_CRUD_SPY/controllers:CargosController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_CRUD_SPY/controllers:ComentariosController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_CRUD_SPY/controllers:ComentariosController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_CRUD_SPY/controllers:ComentariosController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_CRUD_SPY/controllers:ComentariosController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_CRUD_SPY/controllers:ComentariosController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_CRUD_SPY/controllers:ComentariosController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_CRUD_SPY/controllers:ComentariosController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_CRUD_SPY/controllers:ComentariosController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_CRUD_SPY/controllers:ComentariosController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_CRUD_SPY/controllers:ComentariosController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_CRUD_SPY/controllers:CredencialesController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_CRUD_SPY/controllers:CredencialesController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_CRUD_SPY/controllers:CredencialesController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_CRUD_SPY/controllers:CredencialesController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_CRUD_SPY/controllers:CredencialesController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_CRUD_SPY/controllers:CredencialesController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_CRUD_SPY/controllers:CredencialesController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_CRUD_SPY/controllers:CredencialesController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_CRUD_SPY/controllers:CredencialesController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_CRUD_SPY/controllers:CredencialesController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_CRUD_SPY/controllers:EstacionamientosController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_CRUD_SPY/controllers:EstacionamientosController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_CRUD_SPY/controllers:EstacionamientosController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_CRUD_SPY/controllers:EstacionamientosController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_CRUD_SPY/controllers:EstacionamientosController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_CRUD_SPY/controllers:EstacionamientosController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_CRUD_SPY/controllers:EstacionamientosController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_CRUD_SPY/controllers:EstacionamientosController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_CRUD_SPY/controllers:EstacionamientosController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_CRUD_SPY/controllers:EstacionamientosController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_CRUD_SPY/controllers:EstacionamientosPromocionesController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_CRUD_SPY/controllers:EstacionamientosPromocionesController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_CRUD_SPY/controllers:EstacionamientosPromocionesController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_CRUD_SPY/controllers:EstacionamientosPromocionesController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_CRUD_SPY/controllers:EstacionamientosPromocionesController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_CRUD_SPY/controllers:EstacionamientosPromocionesController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_CRUD_SPY/controllers:EstacionamientosPromocionesController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_CRUD_SPY/controllers:EstacionamientosPromocionesController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_CRUD_SPY/controllers:EstacionamientosPromocionesController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_CRUD_SPY/controllers:EstacionamientosPromocionesController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_CRUD_SPY/controllers:PagosController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_CRUD_SPY/controllers:PagosController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_CRUD_SPY/controllers:PagosController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_CRUD_SPY/controllers:PagosController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_CRUD_SPY/controllers:PagosController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_CRUD_SPY/controllers:PagosController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_CRUD_SPY/controllers:PagosController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_CRUD_SPY/controllers:PagosController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_CRUD_SPY/controllers:PagosController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_CRUD_SPY/controllers:PagosController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_CRUD_SPY/controllers:PromocionesController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_CRUD_SPY/controllers:PromocionesController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_CRUD_SPY/controllers:PromocionesController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_CRUD_SPY/controllers:PromocionesController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_CRUD_SPY/controllers:PromocionesController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_CRUD_SPY/controllers:PromocionesController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_CRUD_SPY/controllers:PromocionesController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_CRUD_SPY/controllers:PromocionesController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_CRUD_SPY/controllers:PromocionesController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_CRUD_SPY/controllers:PromocionesController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_CRUD_SPY/controllers:RankingController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_CRUD_SPY/controllers:RankingController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_CRUD_SPY/controllers:RankingController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_CRUD_SPY/controllers:RankingController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_CRUD_SPY/controllers:RankingController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_CRUD_SPY/controllers:RankingController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_CRUD_SPY/controllers:RankingController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_CRUD_SPY/controllers:RankingController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_CRUD_SPY/controllers:RankingController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_CRUD_SPY/controllers:RankingController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_CRUD_SPY/controllers:RolesController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_CRUD_SPY/controllers:RolesController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_CRUD_SPY/controllers:RolesController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_CRUD_SPY/controllers:RolesController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_CRUD_SPY/controllers:RolesController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_CRUD_SPY/controllers:RolesController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_CRUD_SPY/controllers:RolesController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_CRUD_SPY/controllers:RolesController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_CRUD_SPY/controllers:RolesController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_CRUD_SPY/controllers:RolesController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_CRUD_SPY/controllers:RolesUsuarioController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_CRUD_SPY/controllers:RolesUsuarioController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_CRUD_SPY/controllers:RolesUsuarioController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_CRUD_SPY/controllers:RolesUsuarioController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_CRUD_SPY/controllers:RolesUsuarioController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_CRUD_SPY/controllers:RolesUsuarioController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_CRUD_SPY/controllers:RolesUsuarioController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_CRUD_SPY/controllers:RolesUsuarioController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_CRUD_SPY/controllers:RolesUsuarioController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_CRUD_SPY/controllers:RolesUsuarioController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_CRUD_SPY/controllers:SlotsController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_CRUD_SPY/controllers:SlotsController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_CRUD_SPY/controllers:SlotsController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_CRUD_SPY/controllers:SlotsController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_CRUD_SPY/controllers:SlotsController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_CRUD_SPY/controllers:SlotsController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_CRUD_SPY/controllers:SlotsController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_CRUD_SPY/controllers:SlotsController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_CRUD_SPY/controllers:SlotsController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_CRUD_SPY/controllers:SlotsController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_CRUD_SPY/controllers:TipoPagosController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_CRUD_SPY/controllers:TipoPagosController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_CRUD_SPY/controllers:TipoPagosController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_CRUD_SPY/controllers:TipoPagosController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_CRUD_SPY/controllers:TipoPagosController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_CRUD_SPY/controllers:TipoPagosController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_CRUD_SPY/controllers:TipoPagosController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_CRUD_SPY/controllers:TipoPagosController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_CRUD_SPY/controllers:TipoPagosController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_CRUD_SPY/controllers:TipoPagosController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_CRUD_SPY/controllers:UsuariosController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_CRUD_SPY/controllers:UsuariosController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_CRUD_SPY/controllers:UsuariosController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_CRUD_SPY/controllers:UsuariosController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_CRUD_SPY/controllers:UsuariosController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_CRUD_SPY/controllers:UsuariosController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_CRUD_SPY/controllers:UsuariosController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_CRUD_SPY/controllers:UsuariosController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_CRUD_SPY/controllers:UsuariosController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_CRUD_SPY/controllers:UsuariosController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_CRUD_SPY/controllers:VehiculosController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_CRUD_SPY/controllers:VehiculosController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_CRUD_SPY/controllers:VehiculosController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_CRUD_SPY/controllers:VehiculosController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_CRUD_SPY/controllers:VehiculosController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_CRUD_SPY/controllers:VehiculosController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_CRUD_SPY/controllers:VehiculosController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_CRUD_SPY/controllers:VehiculosController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_CRUD_SPY/controllers:VehiculosController"] = append(beego.GlobalControllerRouter["github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_CRUD_SPY/controllers:VehiculosController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
