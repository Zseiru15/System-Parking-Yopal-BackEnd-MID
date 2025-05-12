package controllers

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_MID_SPY/security"
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

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &body_ingresa); err != nil {
		fmt.Println("Error al procesar el cuerpo de la solicitud:", err)
		c.CustomAbort(400, "Cuerpo de la solicitud inválido")
		return
	}

	jsoncontasena := map[string]interface{}{
		"Contrasena": body_ingresa["password"],
	}
	json_contrasena_byte, _ := json.Marshal(jsoncontasena)

	body_contrasena_byte, _ := services.Metodo_post("CRUD_SPY", "credenciales", json_contrasena_byte)

	fmt.Println("Respuesta del servicio:", string(body_contrasena_byte))

	body_contrasena_json, _ := services.ProcesarJson(body_contrasena_byte)
	Id_contrasena := body_contrasena_json["data"].(map[string]interface{})["Id"]
	fmt.Println("Id de la contraseña:", Id_contrasena)
	Id_string := fmt.Sprintf("%v", Id_contrasena)
	Id_contrasena_int, _ := strconv.Atoi(Id_string)

	rol := body_ingresa["type"].(string)

	Id_rol := services.ObtenerIDRol(rol)
	phone := body_ingresa["phone"]
	phone_string := fmt.Sprintf("%v", phone)
	phone_float, _ := strconv.ParseFloat(phone_string, 64)

	json_usuario := map[string]interface{}{
		"Nombres":                      body_ingresa["firstName"],
		"Apellidos":                    body_ingresa["lastName"],
		"NumeroIdentificacionUsuarios": body_ingresa["documentNumber"],
		"Telefono":                     phone_float,
		"Email":                        body_ingresa["email"],
		"IdContrasenaFk":               map[string]interface{}{"Id": Id_contrasena_int},
		"IdRolesFk":                    map[string]interface{}{"Id": Id_rol},
	}

	json_usuario_byte, _ := json.Marshal(json_usuario)
	response_usuario, err := services.Metodo_post("CRUD_SPY", "usuarios", json_usuario_byte)
	if err != nil {
		fmt.Println("Error al registrar el usuario:", err)
		c.CustomAbort(500, "Error interno al registrar usuario")
		return
	}
	fmt.Println("Respuesta del servicio:", string(response_usuario))

	token, err := security.GenerarToken(fmt.Sprintf("%v", body_ingresa["email"]))
	if err != nil {
		c.CustomAbort(500, "Error generando token")
		return
	}

	// Respuesta JSON optimizada
	c.Data["json"] = map[string]interface{}{
		"Success": true,
		"Status":  201,
		"type":    "post",
		"Message": "Creacion existosa",
		"Data":    json_usuario,
		"token":   token,
	}

	c.Ctx.Output.ContentType("application/json") // Opcional pero recomendado
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
	// Obtener todos los usuarios desde el servicio externo
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
			"Foto":            user["IdUsuariosFk"].(map[string]interface{})["Imagen"],
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
