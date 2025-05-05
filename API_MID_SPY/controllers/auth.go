package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_MID_SPY/services"
)

// AuthController operations for Auth
type AuthController struct {
	beego.Controller
}

// URLMapping ...
func (c *AuthController) URLMapping() {
	c.Mapping("Post", c.Register)
	c.Mapping("Login", c.Login)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Register ...
// @Title Create
// @Description create Usuarios
// @Param	body		body 	models.Usuarios	true		"body for Usuarios content"
// @Success 201 {object} models.Usuarios
// @Failure 403 body is empty
// @router / [post]
func (c *AuthController) Register() {
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
	json_usuario_byte, err := json.Marshal(body_ingresa)
	if err != nil {
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  500,
			"Message": "Error al convertir datos a JSON",
		}
		c.ServeJSON()
		return
	}

	// Llamar a Metodo_post para crear el usuario en CRUD_SPY
	response_usuario, err := services.Metodo_post("CRUD_SPY", "Usuarios", json_usuario_byte)
	if err != nil {
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  500,
			"Message": "Error al crear el usuario",
		}
		c.ServeJSON()
		return
	}

	// Decodificar la respuesta del servicio
	var temporal_usuario2 map[string]interface{}
	if err := json.Unmarshal(response_usuario, &temporal_usuario2); err != nil {
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  500,
			"Message": "Error al procesar la respuesta del servidor",
		}
		c.ServeJSON()
		return
	}

	// 📌 IMPRIMIR LA RESPUESTA COMPLETA
	fmt.Println("Respuesta del servidor CRUD_SPY:", temporal_usuario2)

	// Validar si "data" está presente en la respuesta
	data, ok := temporal_usuario2["data"].(map[string]interface{})
	if !ok {
		c.Data["json"] = map[string]interface{}{
			"Success":     false,
			"Status":      500,
			"Message":     "Estructura de datos incorrecta en la respuesta del servidor",
			"RawResponse": temporal_usuario2, // 🔍 Incluir la respuesta completa para depuración
		}
		c.ServeJSON()
		return
	}

	// Respuesta JSON optimizada
	c.Data["json"] = map[string]interface{}{
		"Succes":  true,
		"Status":  201,
		"type":    "post",
		"Message": "Creacion existosa",
		"Data":    data,
	}

	c.ServeJSON()
}

// Login ...
// @Title Login
// @Description Login Auth by Email
// @Param	Email		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Usuarios
// @Failure 403 :Email is empty
// @router /:Email [get]
func (c *AuthController) Login() {
	id_ingreso := c.Ctx.Input.Param(":Email") // para capturar el parametro del url /id
	// println("ese es el id: ", id_ingreso)
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
		"Id":              userData["data"].(map[string]interface{})["Id"],
		"Email":                userData["data"].(map[string]interface{})["Email"],
		"Contraseña":            userData["data"].(map[string]interface{})["IdContrasenaFk"].(map[string]interface{})["Id"],
		"Rol":            userData["data"].(map[string]interface{})["IdRolesFk"].(map[string]interface{})["Id"],
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
// @Description get Auth
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Auth
// @Failure 403
// @router / [get]
func (c *AuthController) GetAll() {

}

// Put ...
// @Title Put
// @Description update the Auth
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Auth	true		"body for Auth content"
// @Success 200 {object} models.Auth
// @Failure 403 :id is not int
// @router /:id [put]
func (c *AuthController) Put() {

}

// Delete ...
// @Title Delete
// @Description delete the Auth
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *AuthController) Delete() {

}
