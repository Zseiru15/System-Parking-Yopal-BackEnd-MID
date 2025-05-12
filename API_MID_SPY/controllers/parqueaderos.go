package controllers

import (
	"encoding/json"
	"fmt"

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

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &body_ingresa); err != nil {
		fmt.Println("Error al procesar el cuerpo de la solicitud:", err)
		c.CustomAbort(400, "Cuerpo de la solicitud inválido")
		return
	}

	// Validaciones básicas
	requiredFields := []string{"parkingName", "latitude", "length", "address"}
	for _, field := range requiredFields {
		if _, ok := body_ingresa[field]; !ok {
			c.CustomAbort(400, fmt.Sprintf("Campo obligatorio faltante: %s", field))
			return
		}
	}

	json_parqueadero := map[string]interface{}{
		"Nombres":    body_ingresa["parkingName"],
		"Latitud":    body_ingresa["latitude"],
		"Longitud":   body_ingresa["length"],
		"Direccion":  body_ingresa["address"],
		"Carros":     body_ingresa["cars"],
		"Motos":      body_ingresa["motorcycles"],
		"Bicicletas": body_ingresa["bicycles"],
		"Altura":     body_ingresa["height"],
		"Tipo":       body_ingresa["type"],
		"Pisos":      body_ingresa["floor"],
		"Sombra":     body_ingresa["shade"],
	}

	json_parqueadero_byte, err := json.Marshal(json_parqueadero)
	if err != nil {
		fmt.Println("Error al convertir datos a JSON:", err)
		c.CustomAbort(500, "Error interno al procesar datos del parqueadero")
		return
	}

	response_parqueadero, err := services.Metodo_post("CRUD_SPY", "parqueaderos", json_parqueadero_byte)
	if err != nil {
		fmt.Println("Error al registrar el parqueadero:", err)
		c.CustomAbort(500, "Error interno al registrar parqueadero")
		return
	}

	fmt.Println("Respuesta del servicio:", string(response_parqueadero))

	// Respuesta exitosa
	c.Ctx.Output.SetStatus(201)
	c.Data["json"] = map[string]interface{}{
		"Success": true,
		"Status":  201,
		"type":    "post",
		"Message": "Parqueadero creado exitosamente",
		"Data":    json_parqueadero,
	}
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

	body, err := services.Metodo_get("CRUD_SPY", "Parqueaderos", id_ingreso)
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
	body, err := services.Metodo_get("CRUD_SPY", "Parqueaderos", "")
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
