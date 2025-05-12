package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_MID_SPY/services"
)

// ComentariosController operations for Comentarios
type ComentariosController struct {
	beego.Controller
}

// URLMapping ...
func (c *ComentariosController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Create
// @Description create Comentarios
// @Param	body		body 	models.Comentarios	true		"body for Comentarios content"
// @Success 201 {object} models.Comentarios
// @Failure 403 body is empty
// @router / [post]
func (c *ComentariosController) Post() {

}

// GetOne ...
// @Title GetOne
// @Description get Comentarios by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Comentarios
// @Failure 403 :id is empty
// @router /:id [get]
func (c *ComentariosController) GetOne() {

}

// GetAll ...
// @Title GetAll
// @Description get Comentarios
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Comentarios
// @Failure 403
// @router / [get]
func (c *ComentariosController) GetAll() {
	// Obtener todos los usuarios desde el servicio externo
	body, err := services.Metodo_get("CRUD_SPY", "Comentarios", "")
	if err != nil || len(body) == 0 {
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  500,
			"Message": "Error al obtener los comentarios",
		}
		c.ServeJSON()
		return
	}

	// Decodificar JSON
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

	// Validar si "data" contiene usuarios
	commentsArray, ok := responseData["data"].([]interface{})
	if !ok {
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  500,
			"Message": "Estructura de datos incorrecta",
		}
		c.ServeJSON()
		return
	}

	// Procesar cada usuario en la lista
	var comentarios []map[string]interface{}
	for _, item := range commentsArray {
		user, ok := item.(map[string]interface{})
		if !ok {
			continue
		}

		// Agregar usuario procesado a la lista final
		comentarios = append(comentarios, map[string]interface{}{
			"Foto":             user["IdUsuariosFk"].(map[string]interface{})["Imagen"],
			"Nombres":         user["IdUsuariosFk"].(map[string]interface{})["Nombres"],
			"Apellidos":       user["IdUsuariosFk"].(map[string]interface{})["Apellidos"],
			"Vehiculo":        user["IdVehiculosFk"].(map[string]interface{})["Marca"],
			"Estacionamiento": user["IdEstacionamientoFk"].(map[string]interface{})["Nombres"],
			"Comentario":      user["Comentario"],
			"Calificacion":    user["Calificacion"],
			"Fecha":           user["FechaRegistro"],
		})
	}

	// Respuesta JSON optimizada
	c.Data["json"] = map[string]interface{}{
		"Success": true,
		"Status":  200,
		"Message": "Consulta exitosa",
		"Data":    comentarios,
		"Total":   len(comentarios), // Indica cuántos usuarios se obtuvieron
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the Comentarios
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Comentarios	true		"body for Comentarios content"
// @Success 200 {object} models.Comentarios
// @Failure 403 :id is not int
// @router /:id [put]
func (c *ComentariosController) Put() {

}

// Delete ...
// @Title Delete
// @Description delete the Comentarios
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *ComentariosController) Delete() {

}
