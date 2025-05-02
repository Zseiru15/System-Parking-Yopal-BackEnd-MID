package models

import (
	"time"

)

// Modelo Credenciales (debe estar definido)
type Credenciales struct {
	Id             int       `orm:"column(Id_Credenciales);pk;auto"`
	HashContrasena string    `orm:"column(Contrasena);size(255)"`
	Estado         bool      `orm:"column(Estado);default(true)"`
	FechaRegistro  time.Time `orm:"column(Fecha_Registro);auto_now_add;type(timestamp)"`
	FechaModifica  time.Time `orm:"column(Fecha_Modifica);auto_now;type(timestamp)"`
}