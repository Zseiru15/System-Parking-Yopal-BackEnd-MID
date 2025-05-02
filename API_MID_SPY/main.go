package main

import (
	_ "github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_MID_SPY/routers"
	
	_ "github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_MID_SPY/models" // importa aquí para que se registre el modelo

	"github.com/astaxie/beego"

	"github.com/astaxie/beego/orm"
    _ "github.com/lib/pq" // el driver de PostgreSQL (cámbialo si usas otro)
)

func init() {
    // Configura tu base de datos
    orm.RegisterDriver("postgres", orm.DRPostgres)
    orm.RegisterDataBase("default", "postgres", "user=postgres password=posgres dbname=SPY host=localhost sslmode=disable")
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	beego.Run()
}
