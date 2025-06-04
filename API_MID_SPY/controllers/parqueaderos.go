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

	// Obtener ID de usuario de forma segura
	var usuarioId interface{}
	if idMap, ok := body_ingresa["IdAdministradoresFk"].(map[string]interface{}); ok {
		usuarioId = map[string]interface{}{"Id": idMap["Id"]}
	} else {
		c.CustomAbort(400, "Formato inválido para IdAdministradoresFk")
		return
	}

	number := body_ingresa["number"]
	number_string := fmt.Sprintf("%v", number)
	number_float, _ := strconv.ParseFloat(number_string, 64)

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

	sombraStr := fmt.Sprintf("%v", body_ingresa["shade"])
	shadebool := services.ObtenerSombraBoolean(sombraStr)

	// Armado del JSON para enviar al CRUD
	json_parqueadero := map[string]interface{}{
		"IdAdministradoresFk": usuarioId,
		"Nombres":             body_ingresa["parkingName"],
		"Telefono":            number_float,
		"Latitud":             latitude_float,
		"Longitud":            length_float,
		"Email":               body_ingresa["email"],
		"Direccion":           body_ingresa["address"],
		"Carros":              cars_float,
		"Motos":               motorcycles_float,
		"Bicicletas":          bicycles_float,
		"Largo":               body_ingresa["long"],
		"Ancho":               body_ingresa["broad"],
		"Altura":              body_ingresa["height"],
		"Tipo":                body_ingresa["type"],
		"Pisos":               floor_float,
		"Sombra":              shadebool,
		"Descripcion":         body_ingresa["description"],
		"Imagen":              body_ingresa["Imagen"],
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

// GetByUsuario ...
// @Title GetByUsuario
// @Description Obtener parqueaderos asociados a un usuario
// @Param	id		path 	string	true		"ID del usuario"
// @Success 200 {object} []models.Vehiculos
// @Failure 400 El ID es inválido
// @Failure 500 Error interno del servidor
// @router /usuario/:id [get]
func (c *ParqueaderosController) GetByUsuario() {
	idUsuario := c.Ctx.Input.Param(":id")

	if idUsuario == "" {
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  400,
			"Message": "ID de usuario no proporcionado",
		}
		c.ServeJSON()
		return
	}

	// 🚨 Importante: no dejes "/" al final del query
	query := "?query=IdAdministradoresFk.Id:" + idUsuario + "&limit=0"

	// Query param: ?query=IdAdministradorFk.Id:ID
	body, err := services.Metodo_get("CRUD_SPY", "parqueaderos", query)
	if err != nil || len(body) == 0 {
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  404,
			"Message": "No se encontraron parqueaderos para este usuario",
		}
		c.ServeJSON()
		return
	}

	var respuesta map[string]interface{}
	if err := json.Unmarshal(body, &respuesta); err != nil {
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  500,
			"Message": "Error al procesar la respuesta del servidor",
		}
		c.ServeJSON()
		return
	}

	c.Data["json"] = map[string]interface{}{
		"Success": true,
		"Status":  200,
		"Message": "Consulta exitosa",
		"Data":    respuesta["data"],
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

// Put ...
// @Title Update
// @Description update Parqueaderos
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body	body 	models.Parqueaderos	true		"body for Parqueaderos content"
// @Success 200 {object} models.Parqueaderos
// @Failure 400 Invalid ID
// @Failure 500 Error updating user
// @router /:id [put]
func (c *ParqueaderosController) Put() {
	id_ingreso := c.Ctx.Input.Param(":id")

	// Obtener datos actuales del vehiculo desde el CRUD
	actualDataRaw, err := services.Metodo_get("CRUD_SPY", "parqueaderos", id_ingreso)
	if err != nil || len(actualDataRaw) == 0 {
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  404,
			"Message": "No se pudo obtener el parqueadero actual",
		}
		c.ServeJSON()
		return
	}

	var actualParsed map[string]interface{}
	if err := json.Unmarshal(actualDataRaw, &actualParsed); err != nil {
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  500,
			"Message": "Error al leer parqueadero actual",
		}
		c.ServeJSON()
		return
	}

	actualUser := actualParsed["data"].(map[string]interface{})

	// Leer el nuevo cuerpo que se quiere actualizar
	var body_actualizacion map[string]interface{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &body_actualizacion); err != nil {
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  400,
			"Message": "Error al procesar el cuerpo de la solicitud",
		}
		c.ServeJSON()
		return
	}

	// Actualizar solo los campos enviados (no sobrescribir vacíos)
	for key, value := range body_actualizacion {
		actualUser[key] = value
	}

	// Convertir a JSON para el PUT
	json_usuario_byte, err := json.Marshal(actualUser)
	if err != nil {
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  500,
			"Message": "Error al convertir datos a JSON",
		}
		c.ServeJSON()
		return
	}

	// Enviar al CRUD_SPY
	response_usuario, err := services.Metodo_put("CRUD_SPY", "parqueaderos", id_ingreso, json_usuario_byte)
	if err != nil {
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  500,
			"Message": "Error al actualizar el parqueaderos",
		}
		c.ServeJSON()
		return
	}

	var result map[string]interface{}
	if err := json.Unmarshal(response_usuario, &result); err != nil {
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  500,
			"Message": "Error al procesar la respuesta del CRUD",
		}
		c.ServeJSON()
		return
	}

	c.Data["json"] = map[string]interface{}{
		"Success": true,
		"Status":  200,
		"type":    "put",
		"Message": "Parqueadero actualizado correctamente",
		"Data":    result["data"],
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

func (c *ParqueaderosController) DesactivarParqueadero() {
	id := c.Ctx.Input.Param(":id")
	if id == "" {
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  400,
			"Message": "ID no proporcionado",
		}
		c.ServeJSON()
		return
	}

	// Enviar petición PUT al CRUD
	_, err := services.Metodo_put("CRUD_SPY", "/parqueaderos/desactivar", id, nil)
	if err != nil {
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  500,
			"Message": "Error al desactivar parqueadero",
		}
		c.ServeJSON()
		return
	}

	c.Data["json"] = map[string]interface{}{
		"Success": true,
		"Status":  200,
		"Message": "Parqueadero desactivado correctamente",
	}
	c.ServeJSON()
}