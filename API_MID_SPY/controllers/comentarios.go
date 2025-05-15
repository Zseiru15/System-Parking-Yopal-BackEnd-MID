package controllers

import (
	"encoding/json"
	"fmt"
	"strconv"

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
	var body_ingresa map[string]interface{}

	// Leer y verificar el cuerpo de la solicitud
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &body_ingresa); err != nil {
		c.CustomAbort(400, "Cuerpo de la solicitud inválido: "+err.Error())
		return
	}

	// Validaciones de campos obligatorios
	requiredFields := []string{"parkingId", "comment", "classification"}
	for _, field := range requiredFields {
		if val, ok := body_ingresa[field]; !ok || val == nil || val == "" {
			c.CustomAbort(400, fmt.Sprintf("Campo obligatorio faltante o vacío: %s", field))
			return
		}
	}

	// Convertir clasificación a float64
	classification := fmt.Sprintf("%v", body_ingresa["classification"])
	classification_float, _ := strconv.ParseFloat(classification, 64)

	// Convertir parkingId a int (podría llegar como float o string)
	parkingId := fmt.Sprintf("%v", body_ingresa["parkingId"])
	parkingId_int, _ := strconv.Atoi(parkingId)

	// Construcción del JSON para enviar al CRUD
	json_comentario := map[string]interface{}{
		"IdUsuariosfk":        map[string]interface{}{"Id": 1}, // 🔥 Usuario quemado por ahora
		"IdEstacionamientofk": map[string]interface{}{"Id": parkingId_int},
		"Comentario":          body_ingresa["comment"],
		"Calificacion":        classification_float,
	}

	json_data, err := json.Marshal(json_comentario)
	if err != nil {
		c.CustomAbort(500, "Error al convertir datos del comentario: "+err.Error())
		return
	}

	// Enviar a CRUD
	resp, err := services.Metodo_post("CRUD_SPY", "comentarios", json_data)
	if err != nil {
		c.CustomAbort(500, "Error al registrar el comentario: "+err.Error())
		return
	}

	parsedResponse, err := services.ProcesarJson(resp)
	if err != nil {
		c.CustomAbort(500, "Error al interpretar la respuesta del CRUD")
		return
	}

	// Éxito
	c.Ctx.Output.SetStatus(201)
	c.Data["json"] = map[string]interface{}{
		"Success": true,
		"Status":  201,
		"type":    "post",
		"Message": "Comentario creado exitosamente",
		"Data":    parsedResponse,
	}
	c.ServeJSON()
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
	// Llamada al servicio CRUD
	body, err := services.Metodo_get("CRUD_SPY", "comentarios", "")
	if err != nil || len(body) == 0 {
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  500,
			"Message": "Error al obtener los comentarios",
		}
		c.ServeJSON()
		return
	}

	// Decodificación
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

	// Validación de estructura
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

	var comentarios []map[string]interface{}
	for _, item := range commentsArray {
		user, ok := item.(map[string]interface{})
		if !ok {
			continue
		}

		usuarioFk, _ := user["IdUsuariosFk"].(map[string]interface{})
		vehiculoFk, _ := user["IdVehiculosFk"].(map[string]interface{})
		parqueaderoFk, _ := user["IdEstacionamientoFk"].(map[string]interface{})

		comentarios = append(comentarios, map[string]interface{}{
			"Foto":            safeGet(usuarioFk, "Imagen"),
			"Nombres":         safeGet(usuarioFk, "Nombres"),
			"Apellidos":       safeGet(usuarioFk, "Apellidos"),
			"Vehiculo":        safeGet(vehiculoFk, "Marca"),
			"Estacionamiento": safeGet(parqueaderoFk, "Nombres"),
			"Comentario":      user["Comentario"],
			"Calificacion":    user["Calificacion"],
			"Fecha":           user["FechaRegistro"],
		})
	}

	c.Data["json"] = map[string]interface{}{
		"Success": true,
		"Status":  200,
		"Message": "Consulta exitosa",
		"Data":    comentarios,
		"Total":   len(comentarios),
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

func safeGet(m map[string]interface{}, key string) interface{} {
	if m == nil {
		return ""
	}
	if val, ok := m[key]; ok {
		return val
	}
	return ""
}
