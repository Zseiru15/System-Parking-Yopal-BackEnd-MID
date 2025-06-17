package controllers

import (
	"fmt"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
)

type ActividadController struct {
	beego.Controller
}

type ActividadUsuario struct {
	Fecha       string `json:"Fecha"`
	Accion      string `json:"Accion"`
	Descripcion string `json:"Descripcion"`
}

// GET /v1/actividades-usuario/:id?desde=YYYY-MM-DD&hasta=YYYY-MM-DD
func (c *ActividadController) GetActividadesPorUsuario() {
	idUsuario, err := c.GetInt(":id")
	if err != nil {
		c.CustomAbort(400, "ID inválido")
		return
	}

	desdeStr := c.GetString("desde")
	hastaStr := c.GetString("hasta")

	if desdeStr == "" || hastaStr == "" {
		c.CustomAbort(400, "Parámetros 'desde' y 'hasta' son obligatorios")
		return
	}

	layout := "2006-01-02"
	_, err1 := time.Parse(layout, desdeStr)
	_, err2 := time.Parse(layout, hastaStr)
	if err1 != nil || err2 != nil {
		c.CustomAbort(400, "Fechas mal formateadas. Usa YYYY-MM-DD")
		return
	}

	o := orm.NewOrm()
	var actividades []ActividadUsuario

	// Comentarios realizados
	var comentarios []struct {
		Comentario      string    `orm:"column(Comentario)"`
		Fecha           time.Time `orm:"column(Fecha_Registro)"`
		Estacionamiento string    `orm:"column(nombre)"`
	}
	_, _ = o.Raw(`
		SELECT c.Comentario, c.Fecha_Registro, e.Nombres as nombre
		FROM comentarios c
		INNER JOIN estacionamientos e ON c.Id_Estacionamiento_fk = e.Id_estacionamientos
		WHERE c.Id_Usuarios_fk = ? AND DATE(c.Fecha_Registro) BETWEEN ? AND ?
	`, idUsuario, desdeStr, hastaStr).QueryRows(&comentarios)

	for _, cmt := range comentarios {
		actividades = append(actividades, ActividadUsuario{
			Fecha:       cmt.Fecha.Format("2006-01-02"),
			Accion:      "Comentario",
			Descripcion: fmt.Sprintf("Comentó en '%s': %s", cmt.Estacionamiento, cmt.Comentario),
		})
	}

	// Pagos realizados
	var pagos []struct {
		Monto    float64   `orm:"column(Amount)"`
		Moneda   string    `orm:"column(Currency)"`
		Estado   bool      `orm:"column(Status)"`
		TipoPago string    `orm:"column(Tipo_Pago)"`
		Fecha    time.Time `orm:"column(Fecha_Pago)"`
	}
	_, _ = o.Raw(`
		SELECT Amount, Currency, Status, Tipo_Pago, Fecha_Pago
		FROM pagos
		WHERE Id_Usuarios_fk = ? AND DATE(Fecha_Pago) BETWEEN ? AND ?
	`, idUsuario, desdeStr, hastaStr).QueryRows(&pagos)

	for _, p := range pagos {
		estado := "Fallido"
		if p.Estado {
			estado = "Completado"
		}
		actividades = append(actividades, ActividadUsuario{
			Fecha:       p.Fecha.Format("2006-01-02"),
			Accion:      "Pago",
			Descripcion: fmt.Sprintf("Pago %.2f %s (%s) - Estado: %s", p.Monto, p.Moneda, p.TipoPago, estado),
		})
	}

	// Vehículos registrados
	var vehiculos []struct {
		Placa string    `orm:"column(Placa)"`
		Fecha time.Time `orm:"column(Fecha_Registro)"`
		Tipo  string    `orm:"column(Tipo)"`
		Marca string    `orm:"column(Marca)"`
	}
	_, _ = o.Raw(`
		SELECT Placa, Fecha_Registro, Tipo, Marca
		FROM vehiculos
		WHERE Id_usuarios_fk = ? AND DATE(Fecha_Registro) BETWEEN ? AND ?
	`, idUsuario, desdeStr, hastaStr).QueryRows(&vehiculos)

	for _, v := range vehiculos {
		actividades = append(actividades, ActividadUsuario{
			Fecha:       v.Fecha.Format("2006-01-02"),
			Accion:      "Vehículo",
			Descripcion: fmt.Sprintf("Registró vehículo: %s (%s %s)", v.Placa, v.Marca, v.Tipo),
		})
	}

	// Parqueaderos registrados
	var parqueaderos []struct {
		Nombre string    `orm:"column(Nombres)"`
		Fecha  time.Time `orm:"column(Fecha_Registro)"`
	}
	_, _ = o.Raw(`
		SELECT Nombres, Fecha_Registro
		FROM estacionamientos
		WHERE Id_Usuarios_fk = ? AND DATE(Fecha_Registro) BETWEEN ? AND ?
	`, idUsuario, desdeStr, hastaStr).QueryRows(&parqueaderos)

	for _, p := range parqueaderos {
		actividades = append(actividades, ActividadUsuario{
			Fecha:       p.Fecha.Format("2006-01-02"),
			Accion:      "Parqueadero",
			Descripcion: fmt.Sprintf("Registró parqueadero: %s", p.Nombre),
		})
	}

	c.Data["json"] = map[string]interface{}{
		"success": true,
		"message": "Actividades del usuario",
		"data":    actividades,
	}
	c.ServeJSON()
}
