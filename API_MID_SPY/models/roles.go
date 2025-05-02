package models

import (
	"time"

)


// Modelo Roles (debe estar definido)
type Roles struct {
	Id            int       `orm:"column(Id_Roles);pk;auto"`
	Roles         string    `orm:"column(Roles);size(50)"`
	Estado        bool      `orm:"column(Estado);default(true)"`
	FechaRegistro time.Time `orm:"column(Fecha_Registro);auto_now_add;type(timestamp)"`
	FechaModifica time.Time `orm:"column(Fecha_Modifica);auto_now;type(timestamp)"`
}
