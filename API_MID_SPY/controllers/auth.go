package controllers

import (
    "github.com/astaxie/beego"
    "github.com/astaxie/beego/orm"
	"golang.org/x/crypto/bcrypt"
    "encoding/json"

    "github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_MID_SPY/models"
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
// POST /auth/login
func (c *AuthController) Login() {
	var reqUsuario models.Usuarios
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &reqUsuario); err != nil {
		c.CustomAbort(400, "Datos inválidos")
		return
	}

	// Buscar usuario por email
	o := orm.NewOrm()
	var usuario models.Usuarios
	err := o.QueryTable("usuario").Filter("Email", reqUsuario.Email).One(&usuario)
	if err == orm.ErrNoRows {
		c.CustomAbort(401, "Correo o contraseña incorrectos")
		return
	}

	// Verificar contraseña
	err = bcrypt.CompareHashAndPassword([]byte(usuario.Contrasena), []byte(reqUsuario.Contrasena))
	if err != nil {
		c.CustomAbort(401, "Correo o contraseña incorrectos")
		return
	}

	c.Data["json"] = map[string]interface{}{
		"message": "Inicio de sesión exitoso",
		"user":    usuario,
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
	var usuario models.Usuarios
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &usuario); err != nil {
		c.CustomAbort(400, "Datos inválidos")
		return
	}

	// Validar que no exista otro usuario con el mismo email o número de identificación
	o := orm.NewOrm()
	exists := models.Usuarios{}
	err := o.QueryTable("usuarios").Filter("Email", usuario.Email).One(&exists)
	if err == nil {
		c.CustomAbort(400, "El email ya está registrado")
		return
	}

	// Encriptar la contraseña
	hash, err := bcrypt.GenerateFromPassword([]byte(usuario.Contrasena), bcrypt.DefaultCost)
	if err != nil {
		c.CustomAbort(500, "Error encriptando contraseña")
		return
	}
	usuario.Contrasena = string(hash)

	// Insertar en la base de datos
	_, err = o.Insert(&usuario)
	if err != nil {
		c.CustomAbort(500, "Error guardando el usuario")
		return
	}

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
