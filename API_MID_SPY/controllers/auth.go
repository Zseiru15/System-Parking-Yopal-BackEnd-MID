package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"net/http"
)

// AuthController operations for Auth
type AuthController struct {
	beego.Controller
}

// URLMapping ...
func (c *AuthController) URLMapping() {
	c.Mapping("Login", c.Login)
	c.Mapping("Register", c.Register)
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Login maneja el inicio de sesión
// @Title Login
// @Description Login de usuario
// @Param	body		body 	map[string]string	true	"Credenciales de login"
// @Success 200 {object} map[string]string
// @Failure 400 el cuerpo es inválido
// @router /login [post]
func (c *AuthController) Login() {
	var credentials map[string]string
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &credentials); err != nil {
		c.Ctx.Output.SetStatus(http.StatusBadRequest)
		c.Data["json"] = map[string]string{"error": "Solicitud inválida"}
		c.ServeJSON()
		return
	}

	email := credentials["Email"]
	password := credentials["password"]

	// Aquí iría la validación real con base de datos
	if email == "admin@example.com" && password == "123456" {
		c.Data["json"] = map[string]string{"message": "Login exitoso"}
	} else {
		c.Ctx.Output.SetStatus(http.StatusUnauthorized)
		c.Data["json"] = map[string]string{"error": "Credenciales inválidas"}
	}
	c.Data["json"] = map[string]interface{}{
		"Success": true,
		"Status":  201,
		"type":    "post",
		"Message": "Creacion exitosa",
		"Data":    c.Data,
	}
	c.ServeJSON()
}

// Register maneja el registro de usuario
// @Title Register
// @Description Registro de usuario
// @Param	body		body 	map[string]string	true	"Datos del usuario"
// @Success 201 {object} map[string]string
// @Failure 400 si los datos son inválidos
// @router /register [post]
func (c *AuthController) Register() {
	var user map[string]string
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &user); err != nil {
		c.Ctx.Output.SetStatus(http.StatusBadRequest)
		c.Data["json"] = map[string]string{"error": "Datos inválidos"}
		c.ServeJSON()
		return
	}

	// Aquí deberías guardar el usuario en la BD (esta es solo una simulación)
	c.Ctx.Output.SetStatus(http.StatusCreated)
	c.Data["json"] = map[string]string{"message": "Usuario registrado exitosamente"}
	c.ServeJSON()
}


// Post ...
// @Title Create
// @Description create Auth
// @Param	body		body 	models.Auth	true		"body for Auth content"
// @Success 201 {object} models.Auth
// @Failure 403 body is empty
// @router / [post]
func (c *AuthController) Post() {

}

// GetOne ...
// @Title GetOne
// @Description get Auth by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Auth
// @Failure 403 :id is empty
// @router /:id [get]
func (c *AuthController) GetOne() {

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
