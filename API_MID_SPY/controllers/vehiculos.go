package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_MID_SPY/services"
)

// VehiculosController operations for Vehiculos
type VehiculosController struct {
	beego.Controller
}

// URLMapping ...
func (c *VehiculosController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Create
// @Description create Vehiculos
// @Param	body		body 	models.Vehiculos	true		"body for Vehiculos content"
// @Success 201 {object} models.Vehiculos
// @Failure 400 Body is empty
// @router / [post]
func (c *VehiculosController) Post() {
	var body_ingresa map[string]interface{}

	// Leer y verificar el cuerpo de la solicitud
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &body_ingresa); err != nil {
		c.CustomAbort(400, "Cuerpo de la solicitud inválido: "+err.Error())
		return
	}

	// Validaciones de campos obligatorios
	requiredFields := []string{"Type", "vehicleBrand", "vehicleModel", "vehicleYear", "vehiclePlate", "IdUsuariosFk"}
	for _, field := range requiredFields {
		if val, ok := body_ingresa[field]; !ok || val == nil || val == "" {
			c.CustomAbort(400, fmt.Sprintf("Campo obligatorio faltante o vacío: %s", field))
			return
		}
	}

	// Obtener ID de usuario de forma segura
	var usuarioId interface{}
	if idMap, ok := body_ingresa["IdUsuariosFk"].(map[string]interface{}); ok {
		usuarioId = idMap["Id"]
	} else {
		c.CustomAbort(400, "Formato inválido para IdUsuariosFk")
		return
	}

	// Armado del JSON para enviar al CRUD
	json_vehiculos := map[string]interface{}{
		"IdUsuariosFk": map[string]interface{}{"Id": usuarioId},
		"Tipo":         body_ingresa["Type"],
		"Marca":        body_ingresa["vehicleBrand"],
		"Modelo":       body_ingresa["vehicleModel"],
		"Año":          body_ingresa["vehicleYear"],
		"Placa":        body_ingresa["vehiclePlate"],
	}

	// Si incluye imagen (opcional)
	if imagen, ok := body_ingresa["Imagen"]; ok {
		json_vehiculos["Imagen"] = imagen
	}

	json_vehiculos_byte, err := json.Marshal(json_vehiculos)
	if err != nil {
		c.CustomAbort(500, "Error al procesar datos del vehículo: "+err.Error())
		return
	}

	// Envío a servicio CRUD
	response_vehiculos, err := services.Metodo_post("CRUD_SPY", "vehiculos", json_vehiculos_byte)
	if err != nil {
		c.CustomAbort(500, "Error al registrar el vehículo: "+err.Error())
		return
	}

	parsedResponse, err := services.ProcesarJson(response_vehiculos)
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
		"Message": "Vehículo creado exitosamente",
		"Data":    parsedResponse,
	}
	c.Ctx.Output.ContentType("application/json")
	c.ServeJSON()
}

// GetOne ...
// @Title GetOne
// @Description get Vehiculos by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Vehiculos
// @Failure 404 Vehicle not found
// @router /:id [get]
func (c *VehiculosController) GetOne() {
	id_ingreso := c.Ctx.Input.Param(":id")

	body, err := services.Metodo_get("CRUD_SPY", "vehiculos", id_ingreso)
	if err != nil || len(body) == 0 {
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  404,
			"Message": "Vehículo no encontrado",
		}
		c.ServeJSON()
		return
	}

	var vehiculoData map[string]interface{}
	if err := json.Unmarshal(body, &vehiculoData); err != nil {
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
		"Data":    vehiculoData["data"],
	}
	c.ServeJSON()
}

// GetByUsuario ...
// @Title GetByUsuario
// @Description Obtener vehículos asociados a un usuario
// @Param	id		path 	string	true		"ID del usuario"
// @Success 200 {object} []models.Vehiculos
// @Failure 400 El ID es inválido
// @Failure 500 Error interno del servidor
// @router /usuario/:id [get]
func (c *VehiculosController) GetByUsuario() {
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
	query := "?query=IdUsuariosFk.Id:" + idUsuario + "&limit=0"

	body, err := services.Metodo_get("CRUD_SPY", "vehiculos", query)
	if err != nil || len(body) == 0 {
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  500,
			"Message": "Error al obtener vehículos del usuario",
		}
		c.ServeJSON()
		return
	}

	var response map[string]interface{}
	if err := json.Unmarshal(body, &response); err != nil {
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  500,
			"Message": "Error al procesar respuesta del servidor",
		}
		c.ServeJSON()
		return
	}

	data, ok := response["data"].([]interface{})
	if !ok {
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  500,
			"Message": "Formato de datos incorrecto en respuesta",
			"Raw":     response,
		}
		c.ServeJSON()
		return
	}

	c.Data["json"] = map[string]interface{}{
		"Success": true,
		"Status":  200,
		"Message": "Vehículos del usuario",
		"Data":    data,
	}
	c.ServeJSON()
}

// GetAll ...
// @Title GetAll
// @Description Obtiene todos los vehículos
// @Success 200 {array} models.Vehiculos
// @Failure 500 Error interno del servidor
// @router / [get]
func (c *VehiculosController) GetAll() {
	body, err := services.Metodo_get("CRUD_SPY", "vehiculos", "")
	if err != nil || len(body) == 0 {
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  500,
			"Message": "Error al obtener los vehículos",
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
// @Description update Vehiculos
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body	body 	models.Vehiculos	true		"body for Vehiculos content"
// @Success 200 {object} models.Vehiculos
// @Failure 400 Invalid ID
// @Failure 500 Error updating user
// @router /:id [put]
func (c *VehiculosController) Put() {
	id_ingreso := c.Ctx.Input.Param(":id")

	// Obtener datos actuales del vehiculo desde el CRUD
	actualDataRaw, err := services.Metodo_get("CRUD_SPY", "vehiculos", id_ingreso)
	if err != nil || len(actualDataRaw) == 0 {
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  404,
			"Message": "No se pudo obtener el vehiculo actual",
		}
		c.ServeJSON()
		return
	}

	var actualParsed map[string]interface{}
	if err := json.Unmarshal(actualDataRaw, &actualParsed); err != nil {
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  500,
			"Message": "Error al leer vehiculo actual",
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
	response_usuario, err := services.Metodo_put("CRUD_SPY", "vehiculos", id_ingreso, json_usuario_byte)
	if err != nil {
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  500,
			"Message": "Error al actualizar el vehiculo",
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
		"Message": "Vehiculo actualizado correctamente",
		"Data":    result["data"],
	}
	c.ServeJSON()
}

// Delete ...
// @Title Disable
// @Description Cambia el estado de un Vehículo a false en lugar de eliminarlo
// @Param	id		path 	string	true		"The ID of the vehicle to disable"
// @Success 200 {string} Vehículo deshabilitado exitosamente
// @Failure 400 ID no válido
// @Failure 404 Vehículo no encontrado
// @router /:id [delete]
func (c *VehiculosController) Delete() {
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
	_, err := services.Metodo_put("CRUD_SPY", "vehiculos", id_ingreso, json_byte)
	if err != nil {
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  500,
			"Message": "Error al deshabilitar el vehículo",
		}
		c.ServeJSON()
		return
	}

	// Confirmar la actualización
	c.Data["json"] = map[string]interface{}{
		"Success": true,
		"Status":  200,
		"Message": "Vehículo deshabilitado correctamente",
	}
	c.ServeJSON()
}

func (c *VehiculosController) DesactivarVehiculo() {
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
	_, err := services.Metodo_put("CRUD_SPY", "/vehiculos/desactivar", id, nil)
	if err != nil {
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  500,
			"Message": "Error al desactivar vehículo",
		}
		c.ServeJSON()
		return
	}

	c.Data["json"] = map[string]interface{}{
		"Success": true,
		"Status":  200,
		"Message": "Vehículo desactivado correctamente",
	}
	c.ServeJSON()
}
