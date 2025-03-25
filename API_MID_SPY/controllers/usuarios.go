package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_MID_SPY/services"
)

// UsuariosController operations for Usuarios
type UsuariosController struct {
	beego.Controller
}

// URLMapping ...
func (c *UsuariosController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Create
// @Description create Usuarios
// @Param	body		body 	models.Usuarios	true		"body for Usuarios content"
// @Success 201 {object} models.Usuarios
// @Failure 403 body is empty
// @router / [post]
func (c *UsuariosController) Post() {
	var body_ingresa []map[string]interface{}
	var alerta models.Alert
	var temporal_usuario []byte
	var temporal_producto []byte

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &body_ingresa); err == nil {
		//fmt.Println("body que ingresa:", body_ingresa)

		jsonData, err := json.MarshalIndent(body_ingresa, "", " ")
		if err != nil {
			fmt.Println("Error al convertir a JSON", err)
		}

		json_usuario := body_ingresa[0]
		json_producto := body_ingresa[1]
		//fmt.Println("Body usuario:", json_usuario)
		//fmt.Println("Body producto:", json_producto)

		json_usuario_byte, _ := json.Marshal(json_usuario)
		response_usuario, _ := services.Metodo_post("Servicio_post", json_usuario_byte)

		temporal_usuario = response_usuario

		//fmt.Println("Esto responde el post de usuario:", string(response_usuario))
		//fmt.Println("Esto respnde el post de producto:", string(response_producto))

		fmt.Print("Body de ingreso en JSON:", string(jsonData))
	}
	var temporal_usuario2 map[string]interface{}
	var temporal_producto2 map[string]interface{}

	json.Unmarshal(temporal_usuario, &temporal_usuario2)
	json.Unmarshal(temporal_producto, &temporal_producto2)
	var body_final []map[string]interface{}
	body_final = append(body_final, temporal_usuario2["data"].(map[string]interface{}))
	body_final = append(body_final, temporal_producto2["data"].(map[string]interface{}))

	alerta.Code = "201"
	alerta.Type = "post"
	alerta.Body = body_final
	c.Data["json"] = alerta
	c.ServeJSON()
}

// GetOne ...
// @Title GetOne
// @Description get Usuarios by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Usuarios
// @Failure 403 :id is empty
// @router /:id [get]
func (c *UsuariosController) GetOne() {
	id_ingreso := c.Ctx.Input.Param(":id") // para capturar el parametro del url /id
	println("ese es el id: ", id_ingreso)
	//----------------------------------------------------------------------------------------
	// println("PASO 1")
	//asignacion de datos al body
	// Obtener datos del usuario desde el servicio externo
	body, err := services.Metodo_get("CRUD_SPY", "Usuarios", id_ingreso)
	if err != nil || len(body) == 0 {
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  500,
			"Message": "Error al obtener el usuario o el usuario no existe",
		}
		c.ServeJSON()
		return
	}
	// println("PASO 1.1")
	//----------------------------------------------------------------------------------------
	// println("PASO 2")
	// Decodificar JSON
	var userData map[string]interface{}
	if err := json.Unmarshal(body, &userData); err != nil {
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  500,
			"Message": "Error al procesar la respuesta del servidor",
		}
		c.ServeJSON()
		return
	}
	// println("PASO 2.1")
	// fmt.Println("Respuesta del servicio:", userData)

	// Extraer y validar solo los datos necesarios
	// El JSON que esperas es un array de objetos
	usuario := map[string]interface{}{
		"Nombres":              userData["data"].(map[string]interface{})["Nombres"],
		"Apellidos":            userData["data"].(map[string]interface{})["Apellidos"],
		"NumeroIdentificacion": userData["data"].(map[string]interface{})["NumeroIdentificacionUsuarios"],
		"Edad":                 userData["data"].(map[string]interface{})["Edad"],
		"Email":                userData["data"].(map[string]interface{})["Email"],
		"Telefono":             userData["data"].(map[string]interface{})["Telefono"],
		"Direccion":            userData["data"].(map[string]interface{})["Direccion"],
		"IdRolesFk":            userData["data"].(map[string]interface{})["IdRolesFk"].(map[string]interface{})["Id"],
	}
	// jsonData2, _ := json.MarshalIndent(usuario, "", "  ")
	// println("PASO 2.2")
	// println("Respuesta del servicio:", string(jsonData2))

	//----------------------------------------------------------------------------------------
	// println("PASO 3")
	// Respuesta JSON optimizada
	c.Data["json"] = map[string]interface{}{
		"Success": true,
		"Status":  200,
		"Message": "Consulta exitosa",
		"Data":    usuario,
	}
	// println("PASO 3.1")
	c.ServeJSON()
}

// GetAll ...
// @Title GetAll
// @Description Obtiene todos los usuarios
// @Success 200 {array} models.Usuarios
// @Failure 500 Error interno del servidor
// @router / [get]
func (c *UsuariosController) GetAll() {
	// Obtener todos los usuarios desde el servicio externo
	body, err := services.Metodo_get("CRUD_SPY", "Usuarios", "")
	if err != nil || len(body) == 0 {
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  500,
			"Message": "Error al obtener los usuarios",
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
	usersArray, ok := responseData["data"].([]interface{})
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
	var usuarios []map[string]interface{}
	for _, item := range usersArray {
		user, ok := item.(map[string]interface{})
		if !ok {
			continue
		}

		// Validar y extraer `IdRolesFk` si existe
		idRolesFk := 0
		if idRoles, ok := user["IdRolesFk"].(map[string]interface{}); ok {
			if id, ok := idRoles["Id"].(float64); ok {
				idRolesFk = int(id)
			}
		}

		// Agregar usuario procesado a la lista final
		usuarios = append(usuarios, map[string]interface{}{
			"Nombres":              user["Nombres"],
			"Apellidos":            user["Apellidos"],
			"NumeroIdentificacion": user["NumeroIdentificacionUsuarios"],
			"Edad":                 user["Edad"],
			"Email":                user["Email"],
			"Telefono":             user["Telefono"],
			"Direccion":            user["Direccion"],
			"IdRolesFk":            idRolesFk,
		})
	}

	// Respuesta JSON optimizada
	c.Data["json"] = map[string]interface{}{
		"Success": true,
		"Status":  200,
		"Message": "Consulta exitosa",
		"Data":    usuarios,
		"Total":   len(usuarios), // Indica cuántos usuarios se obtuvieron
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the Usuarios
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Usuarios	true		"body for Usuarios content"
// @Success 200 {object} models.Usuarios
// @Failure 403 :id is not int
// @router /:id [put]
func (c *UsuariosController) Put() {

}

// Delete ...
// @Title Delete
// @Description delete the Usuarios
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *UsuariosController) Delete() {

}
