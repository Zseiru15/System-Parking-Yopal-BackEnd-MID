package controllers

import (
	"fmt"
	"sort"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
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

	// CONSULTA COMENTARIOS
	var comentarios []struct {
		Comentario      string
		Fecha           time.Time
		Estacionamiento string
	}
	_, err = o.Raw(`
	SELECT 
		c."Comentario" AS Comentario,
		c."Fecha_Registro" AS Fecha,
		e."Nombres" AS Estacionamiento
	FROM "Comentarios" c
	INNER JOIN "Estacionamientos" e ON c."Id_Estacionamiento_fk" = e."Id_estacionamientos"
	WHERE c."Id_Usuarios_fk" = ? AND DATE(c."Fecha_Registro") BETWEEN ? AND ?
`, idUsuario, desdeStr, hastaStr).QueryRows(&comentarios)
	fmt.Println("Comentarios obtenidos:", comentarios)
	if err != nil {
		beego.Error("Error consultando comentarios:", err)
	}
	for _, cmt := range comentarios {
		actividades = append(actividades, ActividadUsuario{
			Fecha:       cmt.Fecha.Format(layout),
			Accion:      "Comentario",
			Descripcion: fmt.Sprintf("Comentó en '%s': %s", cmt.Estacionamiento, cmt.Comentario),
		})
	}

	// CONSULTA PAGOS
	var pagos []struct {
		Monto    float64   `json:"monto"`
		Moneda   string    `json:"moneda"`
		Estado   bool      `json:"estado"`
		TipoPago string    `json:"tipo_pago"`
		Fecha    time.Time `json:"fecha"`
	}
	_, err = o.Raw(`
	SELECT 
		"Amount" AS monto,
		"Currency" AS moneda,
		"Status" AS estado,
		"Tipo_Pago" AS tipo_pago,
		"Fecha_Pago" AS fecha
	FROM "Pagos"
	WHERE "Id_Usuarios_fk" = ? AND DATE("Fecha_Pago") BETWEEN ? AND ?
`, idUsuario, desdeStr, hastaStr).QueryRows(&pagos)
	fmt.Println("Pagos obtenidos:", pagos)
	if err != nil {
		beego.Error("Error consultando pagos:", err)
	}
	for _, p := range pagos {
		estado := "Fallido"
		if p.Estado {
			estado = "Completado"
		}
		actividades = append(actividades, ActividadUsuario{
			Fecha:       p.Fecha.Format(layout),
			Accion:      "Pago",
			Descripcion: fmt.Sprintf("Pago %.2f %s (%s) - Estado: %s", p.Monto, p.Moneda, p.TipoPago, estado),
		})
	}

	// CONSULTA VEHÍCULOS
	var vehiculos []struct {
		Placa string    `json:"placa"`
		Fecha time.Time `json:"fecha"`
		Tipo  string    `json:"tipo"`
		Marca string    `json:"marca"`
	}
	_, err = o.Raw(`
	SELECT 
		"Placa" AS placa,
		"Fecha_Registro" AS fecha,
		"Tipo" AS tipo,
		"Marca" AS marca
	FROM "Vehiculos"
	WHERE "Id_usuarios_fk" = ? AND DATE("Fecha_Registro") BETWEEN ? AND ?
`, idUsuario, desdeStr, hastaStr).QueryRows(&vehiculos)

	fmt.Println("Vehiculos obtenidos:", vehiculos)
	if err != nil {
		beego.Error("Error consultando vehículos:", err)
	}
	for _, v := range vehiculos {
		actividades = append(actividades, ActividadUsuario{
			Fecha:       v.Fecha.Format(layout),
			Accion:      "Vehículo",
			Descripcion: fmt.Sprintf("Registró vehículo: %s (%s %s)", v.Placa, v.Marca, v.Tipo),
		})
	}

	// CONSULTA PARQUEADEROS
	var parqueaderos []struct {
		Nombre string    `json:"nombre"`
		Fecha  time.Time `json:"fecha"`
	}
	_, err = o.Raw(`
	SELECT 
		"Nombres" AS nombre,
		"Fecha_Registro" AS fecha
	FROM "Estacionamientos"
	WHERE "Id_Administradores_fk" = ? AND DATE("Fecha_Registro") BETWEEN ? AND ?
`, idUsuario, desdeStr, hastaStr).QueryRows(&parqueaderos)

	fmt.Println("Parqueaderos obtenidos:", parqueaderos)
	if err != nil {
		beego.Error("Error consultando parqueaderos:", err)
	}
	for _, p := range parqueaderos {
		actividades = append(actividades, ActividadUsuario{
			Fecha:       p.Fecha.Format(layout),
			Accion:      "Parqueadero",
			Descripcion: fmt.Sprintf("Registró parqueadero: %s", p.Nombre),
		})
	}

	// ORDENAR POR FECHA DESCENDENTE
	sort.SliceStable(actividades, func(i, j int) bool {
		ti, _ := time.Parse(layout, actividades[i].Fecha)
		tj, _ := time.Parse(layout, actividades[j].Fecha)
		return ti.After(tj)
	})

	// RESPUESTA
	c.Data["json"] = map[string]interface{}{
		"success": true,
		"message": "Actividades del usuario",
		"data":    actividades,
	}
	c.ServeJSON()
}
