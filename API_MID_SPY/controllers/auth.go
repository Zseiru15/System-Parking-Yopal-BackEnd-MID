package controllers

import (
	"encoding/json"
	"net/http"
	"time"
	
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"golang.org/x/crypto/bcrypt"
	
	"github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_MID_SPY/models"
	"github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_MID_SPY/security"
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

// LoginRequest define la estructura para el login
type LoginRequest struct {
	Email      string `json:"email"`
	Contrasena string `json:"contrasena"`
}

// RegisterRequest define la estructura para registro
type RegisterRequest struct {
	Nombres    string `json:"nombres"`
	Apellidos  string `json:"apellidos"`
	Email      string `json:"email"`
	Contrasena string `json:"contrasena"`
	Telefono   int64  `json:"telefono"`
}

// LoginResponse define la respuesta del login
type LoginResponse struct {
	Token string      `json:"token"`
	User  interface{} `json:"user"`
}

// @Title Login
// @Description Inicio de sesión de usuario
// @Param	body		body 	LoginRequest	true	"Credenciales de usuario"
// @Success 200 {object} LoginResponse
// @Failure 400 Credenciales inválidas
// @router /login [post]
func (c *AuthController) Login() {
	var req LoginRequest
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &req); err != nil {
		c.Data["json"] = map[string]string{"error": "Datos inválidos"}
		c.Ctx.Output.SetStatus(http.StatusBadRequest)
		c.ServeJSON()
		return
	}

	o := orm.NewOrm()
	user := models.Usuarios{Email: req.Email}

	// Cargar usuario con sus relaciones
	if err := o.QueryTable("usuarios").Filter("Email", req.Email).RelatedSel("Credencial", "Rol").One(&user); err != nil {
		c.Data["json"] = map[string]string{"error": "Usuario no encontrado"}
		c.Ctx.Output.SetStatus(http.StatusNotFound)
		c.ServeJSON()
		return
	}

	// Verificar contraseña con el hash almacenado en Credencial
	if user.Credencial == nil || bcrypt.CompareHashAndPassword([]byte(user.Credencial.HashContrasena), []byte(req.Contrasena)) != nil {
		c.Data["json"] = map[string]string{"error": "Credenciales incorrectas"}
		c.Ctx.Output.SetStatus(http.StatusUnauthorized)
		c.ServeJSON()
		return
	}

	token, err := security.GenerateJWT(user)
	if err != nil {
		c.Data["json"] = map[string]string{"error": "Error generando token"}
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		c.ServeJSON()
		return
	}

	// Limpiar datos sensibles
	user.Credencial.HashContrasena = ""
	response := LoginResponse{
		Token: token,
		User:  user,
	}

	c.Data["json"] = response
	c.ServeJSON()
}

// @Title Register
// @Description Registro de nuevo usuario
// @Param	body		body 	RegisterRequest	true	"Datos de usuario"
// @Success 201 {object} models.Usuarios
// @Failure 400 Datos inválidos o usuario ya existe
// @router /register [post]
func (c *AuthController) Register() {
	var req RegisterRequest
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &req); err != nil {
		c.Data["json"] = map[string]string{"error": "Datos inválidos"}
		c.Ctx.Output.SetStatus(http.StatusBadRequest)
		c.ServeJSON()
		return
	}

	o := orm.NewOrm()

	// Iniciar transacción
	to, err := o.Begin()
	if err != nil {
		c.Data["json"] = map[string]string{"error": "Error iniciando transacción"}
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		c.ServeJSON()
		return
	}

	// Verificar si el email ya existe
	if exist := o.QueryTable("usuarios").Filter("Email", req.Email).Exist(); exist {
		to.Rollback()
		c.Data["json"] = map[string]string{"error": "El email ya está registrado"}
		c.Ctx.Output.SetStatus(http.StatusBadRequest)
		c.ServeJSON()
		return
	}

	// 1. Crear credencial primero
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Contrasena), bcrypt.DefaultCost)
	if err != nil {
		to.Rollback()
		c.Data["json"] = map[string]string{"error": "Error procesando contraseña"}
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		c.ServeJSON()
		return
	}

	credencial := models.Credenciales{
		HashContrasena: string(hashedPassword),
		FechaRegistro:  time.Now(),
	}

	idCredencial, err := to.Insert(&credencial)
	if err != nil {
		to.Rollback()
		c.Data["json"] = map[string]string{"error": "Error creando credencial"}
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		c.ServeJSON()
		return
	}

	// 2. Crear usuario
	user := models.Usuarios{
		Nombres:    req.Nombres,
		Apellidos:  req.Apellidos,
		Email:      req.Email,
		Telefono:   req.Telefono,
		Estado:     true,
		Credencial: &models.Credenciales{Id: int(idCredencial)},
		Rol:        &models.Roles{Id: 2}, // Rol de usuario normal
		Usuario:    req.Email,
	}

	_, err = to.Insert(&user)
	if err != nil {
		to.Rollback()
		c.Data["json"] = map[string]string{"error": "Error registrando usuario"}
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		c.ServeJSON()
		return
	}

	// Confirmar transacción
	if err = to.Commit(); err != nil {
		c.Data["json"] = map[string]string{"error": "Error confirmando transacción"}
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		c.ServeJSON()
		return
	}

	// Cargar usuario recién creado con sus relaciones
	o.QueryTable("usuarios").Filter("Id_usuarios", user.IdUsuarios).RelatedSel("Credencial", "Rol").One(&user)

	// Limpiar datos sensibles
	user.Credencial.HashContrasena = ""

	c.Data["json"] = user
	c.Ctx.Output.SetStatus(http.StatusCreated)
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
