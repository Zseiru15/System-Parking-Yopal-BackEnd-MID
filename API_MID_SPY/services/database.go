package services

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
	"log"
)

func init() {
	dbURL := beego.AppConfig.String("sqlconn")
	if dbURL == "" {
		log.Fatal("No se encontró la cadena de conexión 'sqlconn' en app.conf")
	}

	orm.RegisterDriver("postgres", orm.DRPostgres)

	err := orm.RegisterDataBase("default", "postgres", dbURL)
	if err != nil {
		log.Fatalf("Error registrando la base de datos: %v", err)
	}
}
