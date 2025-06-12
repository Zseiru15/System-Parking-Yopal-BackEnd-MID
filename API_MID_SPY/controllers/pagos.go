package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/beego/beego/orm"
	"github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_MID_SPY/models"
	"github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_MID_SPY/services"
)

// PagosController operations for Pagos
type PagosController struct {
	beego.Controller
}

// URLMapping ...
func (c *PagosController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Create
// @Description create Pagos
// @Param	body		body 	models.Pagos	true		"body for Pagos content"
// @Success 201 {object} models.Pagos
// @Failure 403 body is empty
// @router / [post]
func (c *PagosController) Post() {
	var body_ingresa map[string]interface{}

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &body_ingresa); err != nil {
		fmt.Println("Error al procesar el cuerpo de la solicitud:", err)
		c.CustomAbort(400, "Cuerpo de la solicitud inválido")
		return
	}

	// Validación básica
	if body_ingresa["Amount"] == nil || body_ingresa["PayPalOrderID"] == nil {
		c.CustomAbort(400, "Faltan datos obligatorios en el pago")
		return
	}

	// Asignar Status activo
	body_ingresa["Status"] = true

	// Calcular la fecha de expiración (+1 mes)
	now := time.Now()
	fechaFin := now.AddDate(0, 1, 0).Format(time.RFC3339)
	body_ingresa["FechaFin"] = fechaFin

	// Serializar a JSON para enviar al CRUD
	json_pago, err := json.Marshal(body_ingresa)
	if err != nil {
		fmt.Println("Error al serializar el pago:", err)
		c.CustomAbort(500, "Error al preparar datos del pago")
		return
	}

	// Enviar al CRUD
	response_crud, err := services.Metodo_post("CRUD_SPY", "pagos", json_pago)
	if err != nil {
		fmt.Println("Error al enviar el pago al CRUD:", err)
		c.CustomAbort(500, "No se pudo registrar el pago en el CRUD")
		return
	}

	// Procesar respuesta
	var resultado map[string]interface{}
	if err := json.Unmarshal(response_crud, &resultado); err != nil {
		fmt.Println("Error al interpretar respuesta del CRUD:", err)
		c.CustomAbort(500, "Respuesta del CRUD no válida")
		return
	}

	// Respuesta final
	c.Data["json"] = map[string]interface{}{
		"success": true,
		"status":  201,
		"type":    "post",
		"message": "Pago creado correctamente con membresía activa por 1 mes",
		"data":    resultado["data"],
	}
	c.ServeJSON()
}

// GetOne ...
// @Title GetOne
// @Description get Pagos by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Pagos
// @Failure 403 :id is empty
// @router /:id [get]
func (c *PagosController) GetOne() {
	fmt.Println("GetOne Vacio")
}

func GetPayPalAccessToken() (string, error) {
	clientID := os.Getenv("PAYPAL_CLIENT_ID")
	secret := os.Getenv("PAYPAL_SECRET")
	url := "https://api-m.sandbox.paypal.com/v1/oauth2/token" // Cambia a live para producción

	reqBody := strings.NewReader("grant_type=client_credentials")
	req, err := http.NewRequest("POST", url, reqBody)
	if err != nil {
		return "", err
	}

	req.SetBasicAuth(clientID, secret)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return "", errors.New("No se pudo obtener token de PayPal")
	}

	body, _ := ioutil.ReadAll(resp.Body)

	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return "", err
	}

	accessToken, ok := result["access_token"].(string)
	if !ok {
		return "", errors.New("Token inválido")
	}

	return accessToken, nil
}

// Obtener pagos por ID de parqueadero
func (c *PagosController) GetByEstacionamientoId() {
	id := c.Ctx.Input.Param(":id")
	query := fmt.Sprintf("?query=IdEstacionamientosFk:%s&limit=0", id)

	body, err := services.Metodo_get("CRUD_SPY", "pagos", query)
	if err != nil || len(body) == 0 {
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  500,
			"Message": "Error al obtener pagos por parqueadero",
			"error":   err.Error(),
		}
		c.ServeJSON()
		return
	}

	var responseData map[string]interface{}
	if err := json.Unmarshal(body, &responseData); err != nil {
		fmt.Println("RESPUESTA NO JSON:", string(body)) // debug útil
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  500,
			"Message": "Error al procesar la respuesta del servidor",
			"error":   err.Error(),
		}
		c.ServeJSON()
		return
	}

	c.Data["json"] = map[string]interface{}{
		"Success": true,
		"Status":  200,
		"Message": "Consulta exitosa",
		"Data":    responseData["data"],
	}
	c.ServeJSON()
}

// Obtener pagos por ID de usuario
func (c *PagosController) GetByUsuarioId() {
	id := c.Ctx.Input.Param(":id")
	query := fmt.Sprintf("?query=IdUsuariosFk:%s&limit=0", id)

	body, err := services.Metodo_get("CRUD_SPY", "/pagos", query)
	if err != nil || len(body) == 0 {
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  500,
			"Message": "Error al obtener pagos por usuario",
			"error":   err.Error(),
		}
		c.ServeJSON()
		return
	}

	var responseData map[string]interface{}
	if err := json.Unmarshal(body, &responseData); err != nil {
		fmt.Println("RESPUESTA NO JSON:", string(body)) // debug útil
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  500,
			"Message": "Error al procesar la respuesta del servidor",
			"error":   err.Error(),
		}
		c.ServeJSON()
		return
	}

	c.Data["json"] = map[string]interface{}{
		"Success": true,
		"Status":  200,
		"Message": "Consulta exitosa",
		"Data":    responseData["data"],
	}
	c.ServeJSON()
}

// GetAll ...
// @Title GetAll
// @Description get Pagos
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Pagos
// @Failure 403
// @router / [get]
func (c *PagosController) GetAll() {
	body, err := services.Metodo_get("CRUD_SPY", "pagos", "?limit=0")
	if err != nil || len(body) == 0 {
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  500,
			"Message": "Error al obtener los pagos",
			"error":   err.Error(),
		}
		c.ServeJSON()
		return
	}

	var responseData map[string]interface{}
	if err := json.Unmarshal(body, &responseData); err != nil {
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  500,
			"Message": "Error al procesar la respuesta del servidor",
			"error":   err.Error(),
		}
		c.ServeJSON()
		return
	}

	paymentsArray, ok := responseData["data"].([]interface{})
	if !ok {
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  500,
			"Message": "Estructura de datos incorrecta",
		}
		c.ServeJSON()
		return
	}

	var pagos []map[string]interface{}
	for _, item := range paymentsArray {
		payments, ok := item.(map[string]interface{})
		if !ok {
			continue
		}

		// Agregar usuario procesado a la lista final
		pagos = append(pagos, map[string]interface{}{
			"Id":                   payments["Id"],
			"IdUsuariosFk":         payments["IdUsuariosFk"],
			"IdEstacionamientosFk": payments["IdEstacionamientosFk"],
			"PayPalOrdenID":        payments["PayPalOrderID"],
			"Amount":               payments["Amount"],
			"Currency":             payments["Currency"],
			"Status":               payments["Status"],
			"PayerEmail":           payments["PayerEmail"],
			"ReceiverEmail":        payments["ReceiverEmail"],
			"CreatedAt":            payments["CreatedAt"],
		})
	}

	// Respuesta JSON optimizada
	c.Data["json"] = map[string]interface{}{
		"Success": true,
		"Status":  200,
		"Message": "Consulta exitosa",
		"Data":    pagos,
		"Total":   len(pagos), // Indica cuántos usuarios se obtuvieron
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the Pagos
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Pagos	true		"body for Pagos content"
// @Success 200 {object} models.Pagos
// @Failure 403 :id is not int
// @router /:id [put]
func (c *PagosController) Put() {
	c.Ctx.Output.SetStatus(501)
	c.Data["json"] = map[string]interface{}{
		"success": false,
		"status":  501,
		"message": "Método no implementado",
	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the Pagos
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *PagosController) Delete() {
	c.Ctx.Output.SetStatus(501)
	c.Data["json"] = map[string]interface{}{
		"success": false,
		"status":  501,
		"message": "Método no implementado",
	}
	c.ServeJSON()
}

// Endpoint: GET /pagos/membresia-activa/:id_usuario
func (c *PagosController) GetMembresiaActiva() {
	idUsuario := c.Ctx.Input.Param(":id_usuario")
	if idUsuario == "" {
		c.Data["json"] = map[string]interface{}{"error": "ID de usuario requerido"}
		c.Ctx.Output.SetStatus(400)
		c.ServeJSON()
		return
	}

	// Consulta al CRUD por los pagos del usuario
	urlcrud := "http://localhost:8081/v1/pagos?query=IdUsuariosFk:" + idUsuario
	resp, err := http.Get(urlcrud)
	if err != nil {
		c.Data["json"] = map[string]interface{}{"error": "Error al consultar pagos"}
		c.Ctx.Output.SetStatus(500)
		c.ServeJSON()
		return
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		c.Data["json"] = map[string]interface{}{"error": "Error al decodificar respuesta del CRUD"}
		c.Ctx.Output.SetStatus(500)
		c.ServeJSON()
		return
	}

	data, ok := result["Data"].([]interface{})
	if !ok || len(data) == 0 {
		c.Data["json"] = map[string]interface{}{"membresia_activa": false}
		c.ServeJSON()
		return
	}

	// Revisión de fechas
	for _, item := range data {
		pagoMap, ok := item.(map[string]interface{})
		if !ok {
			continue
		}

		status := fmt.Sprintf("%v", pagoMap["Status"])
		if status != "activo" {
			continue
		}

		fechaPagoStr := fmt.Sprintf("%v", pagoMap["CreatedAt"])
		fechaPago, err := time.Parse(time.RFC3339, fechaPagoStr)
		if err != nil {
			continue
		}

		if time.Since(fechaPago) < 30*24*time.Hour {
			c.Data["json"] = map[string]interface{}{"membresia_activa": true}
			c.ServeJSON()
			return
		}
	}

	// Si no encontró membresía activa válida
	c.Data["json"] = map[string]interface{}{"membresia_activa": false}
	c.ServeJSON()
}

// CRON automático: Inactivar membresías vencidas
func InactivarMembresiasVencidas() {
	o := orm.NewOrm()
	var pagos []models.Pagos
	_, err := o.QueryTable("pagos").Filter("Tipo", "membresía").Filter("Status", true).All(&pagos)
	if err != nil {
		fmt.Println("Error al consultar pagos para cron:", err)
		return
	}

	now := time.Now()
	for _, pago := range pagos {
		fechaFin := pago.CreatedAt.AddDate(0, 1, 0) // 1 mes desde creado
		if fechaFin.Before(now) {
			pago.Status = false
			if _, err := o.Update(&pago, "Status"); err != nil {
				fmt.Println("Error al actualizar membresía vencida:", err)
			}
		}
	}
}
