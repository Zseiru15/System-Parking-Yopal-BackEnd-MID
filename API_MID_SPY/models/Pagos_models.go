package models

import (
	"time"

	"github.com/beego/beego/orm"
)

type Pagos struct {
	Id                   int       `json:"Id"`
	IdUsuariosFk         int       `json:"IdUsuariosFk"`
	IdEstacionamientosFk *int      `json:"IdEstacionamientosFk,omitempty"`
	PayPalOrderID        string    `json:"PayPalOrderID"`
	Amount               float64   `json:"Amount"`
	Currency             string    `json:"Currency"`
	PayerEmail           string    `json:"PayerEmail"`
	ReceiverEmail        string    `json:"ReceiverEmail"`
	Tipo                 string    `json:"Tipo"`
	Status               bool      `json:"Status"`
	CreatedAt            time.Time `json:"CreatedAt"`
}

// Función para insertar el pago
func InsertarPago(p *Pagos) error {
	o := orm.NewOrm()
	_, err := o.Insert(&p)
	return err
}
