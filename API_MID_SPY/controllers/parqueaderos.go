package controllers

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_MID_SPY/services"
)

// ParqueaderosController operations for Parqueaderos
type ParqueaderosController struct {
	beego.Controller
}

// URLMapping ...
func (c *ParqueaderosController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Create
// @Description create Parqueaderos
// @Param	body		body 	models.Parqueaderos	true		"body for Parqueaderos content"
// @Success 201 {object} models.Parqueaderos
// @Failure 400 Body is empty
// @router / [post]
func (c *ParqueaderosController) Post() {
	var body_ingresa map[string]interface{}

	// Leer y verificar el cuerpo de la solicitud
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &body_ingresa); err != nil {
		c.CustomAbort(400, "Cuerpo de la solicitud inválido: "+err.Error())
		return
	}

	// Validaciones de campos obligatorios
	requiredFields := []string{"parkingName", "latitude", "length", "address", "height", "type", "floor", "shade"}
	for _, field := range requiredFields {
		if val, ok := body_ingresa[field]; !ok || val == nil || val == "" {
			c.CustomAbort(400, fmt.Sprintf("Campo obligatorio faltante o vacío: %s", field))
			return
		}
	}

	sombraStr := fmt.Sprintf("%v", body_ingresa["shade"])
	shadebool := services.ObtenerSombraBoolean(sombraStr)

	latitude := body_ingresa["latitude"]
	latitude_string := fmt.Sprintf("%v", latitude)
	latitude_float, _ := strconv.ParseFloat(latitude_string, 64)

	length := body_ingresa["length"]
	length_string := fmt.Sprintf("%v", length)
	length_float, _ := strconv.ParseFloat(length_string, 64)

	cars := body_ingresa["cars"]
	cars_string := fmt.Sprintf("%v", cars)
	cars_float, _ := strconv.ParseFloat(cars_string, 64)

	motorcycles := body_ingresa["motorcycles"]
	motorcycles_string := fmt.Sprintf("%v", motorcycles)
	motorcycles_float, _ := strconv.ParseFloat(motorcycles_string, 64)

	bicycles := body_ingresa["bicycles"]
	bicycles_string := fmt.Sprintf("%v", bicycles)
	bicycles_float, _ := strconv.ParseFloat(bicycles_string, 64)

	floor := body_ingresa["floor"]
	floor_string := fmt.Sprintf("%v", floor)
	floor_float, _ := strconv.ParseFloat(floor_string, 64)

	// Armado del JSON para enviar al CRUD
	json_parqueadero := map[string]interface{}{
		"IdAdministradoresFk": map[string]interface{}{"Id": 1}, // ID quemado temporal
		"Nombres":             body_ingresa["parkingName"],
		"Latitud":             latitude_float,
		"Longitud":            length_float,
		"Direccion":           body_ingresa["address"],
		"Carros":              cars_float,
		"Motos":               motorcycles_float,
		"Bicicletas":          bicycles_float,
		"Altura":              body_ingresa["height"],
		"Tipo":                body_ingresa["type"],
		"Pisos":               floor_float,
		"Sombra":              shadebool,
	}

	json_parqueadero_byte, err := json.Marshal(json_parqueadero)
	if err != nil {
		c.CustomAbort(500, "Error al procesar datos del parqueadero: "+err.Error())
		return
	}

	// Envío a servicio CRUD
	response_parqueadero, err := services.Metodo_post("CRUD_SPY", "parqueaderos", json_parqueadero_byte)
	if err != nil {
		c.CustomAbort(500, "Error al registrar el parqueadero: "+err.Error())
		return
	}

	parsedResponse, err := services.ProcesarJson(response_parqueadero)
	if err != nil {
		fmt.Println("Error al procesar la respuesta del CRUD:", err)
		c.CustomAbort(500, "Error al convertir respuesta del CRUD")
		return
	}

	// Éxito
	c.Ctx.Output.SetStatus(201)
	c.Data["json"] = map[string]interface{}{
		"Success": true,
		"Status":  201,
		"type":    "post",
		"Message": "Parqueadero creado exitosamente",
		"Data":    parsedResponse,
	}
	c.Ctx.Output.ContentType("application/json")
	c.ServeJSON()
}

// GetOne ...
// @Title GetOne
// @Description get Parqueaderos by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Parqueaderos
// @Failure 404 Parqueadero not found
// @router /:id [get]
func (c *ParqueaderosController) GetOne() {
	id_ingreso := c.Ctx.Input.Param(":id")

	body, err := services.Metodo_get("CRUD_SPY", "parqueaderos", id_ingreso)
	if err != nil || len(body) == 0 {
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  404,
			"Message": "Parqueadero no encontrado",
		}
		c.ServeJSON()
		return
	}

	var parqueaderoData map[string]interface{}
	if err := json.Unmarshal(body, &parqueaderoData); err != nil {
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  500,
			"Message": "Error al procesar la respuesta del servidor",
		}
		c.ServeJSON()
		return
	}

	// Respuesta JSON optimizada
	c.Data["json"] = map[string]interface{}{
		"Success": true,
		"Status":  200,
		"Message": "Consulta exitosa",
		"Data":    parqueaderoData["data"],
	}
	c.ServeJSON()
}

// GetAll ...
// @Title GetAll
// @Description Obtiene todos los parqueaderos
// @Success 200 {array} models.Parqueaderos
// @Failure 500 Error interno del servidor
// @router / [get]
func (c *ParqueaderosController) GetAll() {
	body, err := services.Metodo_get("CRUD_SPY", "parqueaderos", "")
	if err != nil || len(body) == 0 {
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  500,
			"Message": "Error al obtener los parqueaderos",
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
		}
		c.ServeJSON()
		return
	}

	// Respuesta JSON optimizada
	c.Data["json"] = map[string]interface{}{
		"Success": true,
		"Status":  200,
		"Message": "Consulta exitosa",
		"Data":    responseData["data"],
	}
	c.ServeJSON()
}

// Delete ...
// @Title Disable
// @Description Cambia el estado de un Parqueadero a false en lugar de eliminarlo
// @Param	id		path 	string	true		"The ID of the parking lot to disable"
// @Success 200 {string} Parqueadero deshabilitado exitosamente
// @Failure 400 ID no válido
// @Failure 404 Parqueadero no encontrado
// @router /:id [delete]
func (c *ParqueaderosController) Delete() {
	id_ingreso := c.Ctx.Input.Param(":id")

	if id_ingreso == "" {
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  400,
			"Message": "ID no proporcionado",
		}
		c.ServeJSON()
		return
	}

	// Crear JSON con la actualización del estado
	json_nuevo := map[string]interface{}{
		"Estado": false,
	}

	json_byte, _ := json.Marshal(json_nuevo)

	// Llamar a Metodo_put para actualizar el estado
	_, err := services.Metodo_put("CRUD_SPY", "Parqueaderos", id_ingreso, json_byte)
	if err != nil {
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  500,
			"Message": "Error al deshabilitar el parqueadero",
		}
		c.ServeJSON()
		return
	}

	// Confirmar la actualización
	c.Data["json"] = map[string]interface{}{
		"Success": true,
		"Status":  200,
		"Message": "Parqueadero deshabilitado correctamente",
	}
	c.ServeJSON()
}
