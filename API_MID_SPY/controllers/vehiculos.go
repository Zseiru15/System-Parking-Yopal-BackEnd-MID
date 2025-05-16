package controllers

import (
	"encoding/json"

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

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &body_ingresa); err != nil {
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  400,
			"Message": "Error al procesar el cuerpo de la solicitud",
		}
		c.ServeJSON()
		return
	}

	// Convertir body_ingresa a JSON antes de enviarlo a Metodo_post
	json_vehiculo_byte, err := json.Marshal(body_ingresa)
	if err != nil {
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  500,
			"Message": "Error al convertir datos a JSON",
		}
		c.ServeJSON()
		return
	}

	// Llamar a Metodo_post para crear el vehículo en CRUD_SPY
	response_vehiculo, err := services.Metodo_post("CRUD_SPY", "vehiculos", json_vehiculo_byte)
	if err != nil {
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  500,
			"Message": "Error al registrar el vehículo",
		}
		c.ServeJSON()
		return
	}

	// Decodificar la respuesta del servicio
	var respuesta map[string]interface{}
	if err := json.Unmarshal(response_vehiculo, &respuesta); err != nil {
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  500,
			"Message": "Error al procesar la respuesta del servidor",
		}
		c.ServeJSON()
		return
	}

	// Validar si "data" está presente en la respuesta
	data, ok := respuesta["data"].(map[string]interface{})
	if !ok {
		c.Data["json"] = map[string]interface{}{
			"Success":     false,
			"Status":      500,
			"Message":     "Estructura de datos incorrecta en la respuesta del servidor",
			"RawResponse": respuesta,
		}
		c.ServeJSON()
		return
	}

	// Respuesta JSON optimizada
	c.Data["json"] = map[string]interface{}{
		"Success": true,
		"Status":  201,
		"Type":    "post",
		"Message": "Vehículo registrado exitosamente",
		"Data":    data,
	}
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
// @Description obtiene los vehículos registrados por un usuario específico
// @Param	id		path 	string	true		"ID del usuario"
// @Success 200 {object} []Vehiculos
// @Failure 404 Usuario no tiene vehículos
// @router /usuario/:id [get]
func (c *VehiculosController) GetByUsuario() {
	idUsuario := c.Ctx.Input.Param(":id")

	url := "vehiculos?query=IdUsuariosFk.Id:" + idUsuario

	body, err := services.Metodo_get("CRUD_SPY", url, "")
	if err != nil || len(body) == 0 {
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  404,
			"Message": "No se encontraron vehículos para el usuario",
		}
		c.ServeJSON()
		return
	}

	var responseData map[string]interface{}
	if err := json.Unmarshal(body, &responseData); err != nil {
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  500,
			"Message": "Error al procesar respuesta del servidor",
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
