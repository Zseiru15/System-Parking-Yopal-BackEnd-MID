package main

import (
	"github.com/astaxie/beego/plugins/cors"
	_ "github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_MID_SPY/routers"
	"github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_MID_SPY/services" // ✅ Import correct

	"github.com/astaxie/beego"
)

func main() {

	// Configuración CORS para permitir el acceso desde cualquier origen (entorno de producción)
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Authorization"},
		AllowCredentials: true,
	}))
	//Configuracion desarrollo
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	services.StartScheduler() // ✅ Inicia el scheduler al correr el MID
	beego.Run()
}