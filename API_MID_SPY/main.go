package main

import (
	_ "github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_MID_SPY/routers"
	_ "github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_MID_SPY/models"
	
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
)

func init() {
    orm.RegisterDriver("postgres", orm.DRPostgres)
    
    // Cadena de conexión corregida
    err := orm.RegisterDataBase("default", "postgres", 
        "user=postgres password=posgres dbname=SPY host=localhost port=5432 sslmode=disable")
    
    if err != nil {
        beego.Error("Error al conectar a la base de datos:", err)
        return
    }

    // Forzar recreación de tablas solo en desarrollo
    if beego.BConfig.RunMode == "dev" {
        force := true  // Elimina y recrea las tablas
        verbose := true // Muestra los logs SQL
        err = orm.RunSyncdb("default", force, verbose)
        if err != nil {
            beego.Error("Error al sincronizar la base de datos:", err)
        }
    }
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	beego.Run()
}
