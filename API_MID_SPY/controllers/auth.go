package controllers

import (
	"encoding/json"
	"net/http"
	"github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_MID_SPY/security"
	"github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_MID_SPY/services"
	"github.com/astaxie/beego"
	"golang.org/x/crypto/bcrypt"
)

type AuthController struct {
	beego.Controller
}

// Estructura para respuesta de login
type LoginResponse struct {
	Success bool        `json:"success"`
	Token   string      `json:"token"`
	User    interface{} `json:"user"`
	Message string      `json:"message,omitempty"`
}

// @Title Register
// @Description Registrar nuevo usuario
// @Param	body		body 	models.UserRegister	true	"Datos de registro"
// @Success 201 {object} models.AuthResponse
// @Failure 400 body is empty
// @router /register [post]
func (c *AuthController) Register() {
	var registerData struct {
		Nombres    string `json:"nombres"`
		Apellidos  string `json:"apellidos"`
		Email      string `json:"email"`
		Contrasena string `json:"contrasena"`
	}

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &registerData); err != nil {
		c.Ctx.Output.SetStatus(http.StatusBadRequest)
		c.Data["json"] = map[string]interface{}{
			"success": false,
			"message": "Datos inválidos",
		}
		c.ServeJSON()
		return
	}

	// Hash de la contraseña
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(registerData.Contrasena), bcrypt.DefaultCost)
	if err != nil {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		c.Data["json"] = map[string]interface{}{
			"success": false,
			"message": "Error al procesar la contraseña",
		}
		c.ServeJSON()
		return
	}

	// Crear usuario en CRUD
	userData := map[string]interface{}{
		"Nombres":                      registerData.Nombres,
		"Apellidos":                    registerData.Apellidos,
		"Email":                        registerData.Email,
		"IdContrasenaFk": map[string]interface{}{
			"Contrasena": string(hashedPassword),
			"Estado":     true,
		},
		"Estado":    true,
		"IdRolesFk": 2, // Rol por defecto (2 = usuario normal)
	}

	jsonData, _ := json.Marshal(userData)
	response, err := services.Metodo_post("CRUD_SPY", "Usuarios", jsonData)
	if err != nil {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		c.Data["json"] = map[string]interface{}{
			"success": false,
			"message": "Error al registrar usuario",
		}
		c.ServeJSON()
		return
	}

	c.Data["json"] = map[string]interface{}{
		"success": true,
		"message": "Usuario registrado exitosamente",
		"data":    string(response),
	}
	c.ServeJSON()
}

// @Title Login
// @Description Autenticar usuario
// @Param   body    body    models.LoginRequest  true    "Credenciales"
// @Success 200 {object} models.AuthResponse
// @Failure 401 Unauthorized
// @router /login [post]
func (c *AuthController) Login() {
	// 1. Parsear datos de entrada
	var loginData struct {
		Email     string `json:"email"`
		Password  string `json:"password"`
	}
	
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &loginData); err != nil {
		c.Data["json"] = map[string]interface{}{
			"success": false,
			"message": "Datos inválidos",
		}
		c.Ctx.Output.SetStatus(http.StatusBadRequest)
		c.ServeJSON()
		return
	}

	// 2. Consultar usuario en CRUD
	response, err := services.Metodo_get("CRUD_SPY", "usuarios/by-email", "?email="+loginData.Email)
	if err != nil {
		beego.Error("Error al conectar con CRUD:", err)
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		c.Data["json"] = map[string]interface{}{
			"success": false,
			"message": "Error al conectar con el servidor",
		}
		c.ServeJSON()
		return
	}

	// 3. Parsear respuesta
	var crudResponse struct {
		Success bool                   `json:"success"`
		Data    map[string]interface{} `json:"data"`
	}
	
	if err := json.Unmarshal(response, &crudResponse); !crudResponse.Success || err != nil {
		c.Ctx.Output.SetStatus(http.StatusUnauthorized)
		c.Data["json"] = map[string]interface{}{
			"success": false,
			"message": "Credenciales inválidas",
		}
		c.ServeJSON()
		return
	}

	// 4. Verificar contraseña
	credenciales := crudResponse.Data["IdContrasenaFk"].(map[string]interface{})
	storedPassword := credenciales["Contrasena"].(string)
	
	if err := bcrypt.CompareHashAndPassword(
		[]byte(storedPassword),
		[]byte(loginData.Password),
	); err != nil {
		c.Ctx.Output.SetStatus(http.StatusUnauthorized)
		c.Data["json"] = map[string]interface{}{
			"success": false,
			"message": "Credenciales inválidas",
		}
		c.ServeJSON()
		return
	}

	// 5. Generar token JWT
	userId := int(crudResponse.Data["Id"].(float64))
	token, err := security.GenerateJWT(userId)
	if err != nil {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		c.Data["json"] = map[string]interface{}{
			"success": false,
			"message": "Error al generar token",
		}
		c.ServeJSON()
		return
	}

	// 6. Responder con token
	c.Data["json"] = map[string]interface{}{
		"success": true,
		"token":   token,
		"user": map[string]interface{}{
			"id":    userId,
			"email": crudResponse.Data["Email"],
			"role":  crudResponse.Data["IdRolesFk"].(map[string]interface{})["Id"],
		},
	}
	c.ServeJSON()
}