package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"

	"github.com/astaxie/beego"
	"golang.org/x/crypto/bcrypt"
	"github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_MID_SPY/services"
	jwt "github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_MID_SPY/security"
)

// AuthController operations for Auth
type AuthController struct {
	beego.Controller
}

// URLMapping ...
func (c *AuthController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Orden de las funciones
// 1. Post
// 2. Register (Crear Usuarios)
// 3. Login (Iniciar sesión)
// 4. ValidarToken (Verificar token)
// 5. GeneraryEnviarToken (Generar y enviar token)
// 6. ActualizarContraseña (Actualizar contraseña)
// 7. GetOne
// 8. GetAll
// 9. Put
// 10. GeneraryEnviarToken (Generar y enviar token)
// 11. ActualizarContraseña (Actualizar contraseña)
// 12. Delete
// Funciones auxiliares
// 1. obtenerUsuarioPorCorreo
// 2. obtenerFkCredencialID
// 3. actualizarTokenCredencial
// 4. obtenerCredencialPorID
// 5. desactivarCredencial
// 6. actualizarFkCredencialUsuario
// 7. handleError (manejo de errores)


// @Title Post
// @Description create Auth
// @Param	body		body 	models.Auth	true		"body for Auth content"
// @Success 201 {object} models.Auth
// @Failure 403 body is empty
// @router / [post]
func (c *AuthController) Post() {

}

// Register ...
// @Title Create
// @Description create Usuarios
// @Param	body		body 	models.Usuarios	true		"body for Usuarios content"
// @Success 201 {object} models.Usuarios
// @Failure 403 body is empty
// @router /auth/register [post]
func (c *AuthController) Register() {
	var body_ingreso map[string]interface{}
	var reponseUsuario, responseCredencial, responseRolUsuario []byte

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &body_ingreso); err == nil {
		fmt.Println("Body que ingresa", body_ingreso)

		jsonData, err := json.MarshalIndent(body_ingreso, "", " ")
		if err != nil {
			fmt.Println("Error al convertir a JSON", err)
		}
		fmt.Println("Body de ingreso en JSON:", string(jsonData))

		passStr, _ := body_ingreso["contraseña"].(string)

		// Hashear la contraseña
		hashedPass, err := services.HashContraseña(passStr)
		if err != nil {
			fmt.Println(err)
			c.Data["json"] = map[string]interface{}{"error": "Error interno al procesar la contraseña"}
			c.ServeJSON()
			return
		}

		jsonCredencial := map[string]interface{}{
			"contraseña": hashedPass,
		}
		fmt.Println("este es el json para credenciales: ", jsonCredencial)

		jsonUsuario := map[string]interface{}{
			"nombre":             body_ingreso["Nombre"],
			"apellido":           body_ingreso["Apellido"],
			"contacto":           body_ingreso["Contacto"],
			"correo_electronico": body_ingreso["CorreoElectronico"],
		}
		fmt.Println("este es el json usuario: ", jsonUsuario)

		json_credencial_byte, _ := json.Marshal(jsonCredencial)
		// json_usuario_byte, _ := json.Marshal(jsonUsuario)

		fmt.Println("json credencial: ", string(json_credencial_byte))
		responseCredencial, _ = services.Metodo_post("API_CRUD", "/v1/Credenciales", json_credencial_byte)
		if err != nil {
			fmt.Println("Error al crear credenciales:", err)
			return
		}
		fmt.Println("Respuesta de la API (Credenciales): ", string(responseCredencial))

		// Obtener el ID de credencial creada
		var credencialresponse map[string]interface{}
		if err := json.Unmarshal(responseCredencial, &credencialresponse); err != nil {
			fmt.Println("Error al parsear respuesta de credenciales:", err)
			return
		}

		// Extraer el ID de la credencial
		var credencialID float64
		if data, ok := credencialresponse["Data"].(map[string]interface{}); ok {
			if id, exists := data["Id"].(float64); exists {
				credencialID = id
			} else {
				fmt.Println("Error: No se encontró el ID en la respuesta de Credenciales")
				return
			}
		} else {
			fmt.Println("Error: Estructura de respuesta de credenciales no válida")
			return
		}
		fmt.Println("Id de credenciales: ", credencialID)

		// Crear JSON para Usuario con fk_credencial
		jsonUsuario = map[string]interface{}{
			"Nombre":            body_ingreso["Nombre"],
			"Apellido":          body_ingreso["Apellido"],
			"Contacto":          body_ingreso["Contacto"],
			"CorreoElectronico": body_ingreso["CorreoElectronico"],
			"FkCredencial": map[string]interface{}{
				"Id": credencialID,
			},
		}

		// Convertir a JSON y enviar POST a /v1/Usuario
		json_usuario_byte, _ := json.Marshal(jsonUsuario)
		fmt.Println("Enviando JSON a /v1/Usuario:", string(json_usuario_byte))
		reponseUsuario, err = services.Metodo_post("API_CRUD", "/v1/Usuario", json_usuario_byte)
		if err != nil {
			fmt.Println(" Error al crear usuario:", err)
			return
		}

		fmt.Println("Respuesta de la API (Usuario):", string(reponseUsuario))

		// Extraer el ID del usuario creado
		var usuarioResponse map[string]interface{}
		if err := json.Unmarshal(reponseUsuario, &usuarioResponse); err != nil {
			fmt.Println("Error al parsear respuesta de usuario:", err)
			return
		}

		var usuarioID float64
		if data, ok := usuarioResponse["Data"].(map[string]interface{}); ok {
			if id, exists := data["Id"].(float64); exists {
				usuarioID = id
			} else {
				fmt.Println("Error: No se encontró el ID en la respuesta de Usuario")
				return
			}
		} else {
			fmt.Println("Error: Estructura de respuesta de usuario no válida")
			return
		}

		// Buscar el ID del Rol en la tabla Roles
		rolNombre := body_ingreso["rol"].(string) // Extrae el rol enviado en la solicitud
		var responseRol []byte
		responseRol, err = services.Metodo_get("API_CRUD", "/v1/Roles?query=nombre:", rolNombre)

		if err != nil {
			fmt.Println("Error al obtener el rol:", err)
			return
		}

		var rolResponse map[string]interface{}
		if err := json.Unmarshal(responseRol, &rolResponse); err != nil {
			fmt.Println("Error al parsear respuesta de roles:", err)
			return
		}

		var rolID float64
		if roles, ok := rolResponse["Data"].([]interface{}); ok && len(roles) > 0 {
			if rolData, exists := roles[0].(map[string]interface{}); exists {
				rolID = rolData["Id"].(float64)
			}
		} else {
			fmt.Println("Error: No se encontró el rol en la base de datos")
			return
		}

		// Crear registro en RolesUsuario
		jsonRolUsuario := map[string]interface{}{
			"FkUsuarioRoles": map[string]interface{}{
				"Id": usuarioID,
			},
			"FkRolesUsuario": map[string]interface{}{
				"Id": rolID,
			},
		}

		json_rol_usuario_byte, _ := json.Marshal(jsonRolUsuario)
		responseRolUsuario, _ = services.Metodo_post("API_CRUD", "/v1/Roles_Usuario", json_rol_usuario_byte)

		fmt.Println("Respuesta de la API (RolesUsuario):", string(responseRolUsuario))
	}

	c.Data["json"] = map[string]interface{}{
		"Message": "¡Usuario creado exitosamente!",
	}
	c.ServeJSON()
}

// Post ...
// @Title Login
// @Description create Usuario_controller.Go
// @Param	body		body 	models.Usuario_controller.Go	true		"body for Usuario_controller.Go content"
// @Success 201 {object} models.Usuario_controller.Go
// @Failure 403 body is empty
// @router /auth/login [post]
func (c *AuthController) Login() {
	fmt.Println("Función POST: Login")
	var body map[string]interface{}

	// Leer el cuerpo de la solicitud
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &body); err != nil {
		handleError(c, "Solicitud inválida", err)
		return
	}
	fmt.Println("Datos de login recibidos:", body)

	// Validar que se hayan enviado el correo y la contraseña
	correo, ok := body["CorreoElectronico"].(string)
	if !ok || correo == "" {
		handleError(c, "El campo 'CorreoElectronico' es obligatorio", nil)
		return
	}
	password, ok := body["contraseña"].(string)
	if !ok || password == "" {
		handleError(c, "El campo 'contraseña' es obligatorio", nil)
		return
	}

	// Buscar usuario por correo
	usuario, err := obtenerUsuarioPorCorreo(correo)
	if err != nil {
		handleError(c, "El usuario no existe", err)
		return
	}
	fmt.Println("Usuario encontrado:", usuario)

	usuarioID, ok := usuario["Id"].(float64)
	if !ok {
		handleError(c, "No se pudo obtener el ID del usuario", nil)
		return
	}
	fmt.Println("ID del usuario:", int(usuarioID))

	// Obtener el ID de la credencial asociada
	credencialID, err := obtenerFkCredencialID(usuario)
	if err != nil {
		handleError(c, "No se encontraron credenciales asociadas", err)
		return
	}
	fmt.Println("ID de la credencial:", credencialID)

	// Obtener la credencial (incluyendo la contraseña hasheada)
	credencial, err := obtenerCredencialPorID(credencialID)
	if err != nil {
		handleError(c, "Error al obtener la credencial del usuario", err)
		return
	}
	fmt.Println("Credencial obtenida:", credencial)

	// Comparar la contraseña ingresada con la almacenada
	storedPass, ok := credencial["Contraseña"].(string)
	if !ok || storedPass == "" {
		handleError(c, "La contraseña almacenada es inválida", nil)
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(storedPass), []byte(password))
	if err != nil {
		handleError(c, "Credenciales inválidas", err)
		return
	}
	usuarioIDStr := fmt.Sprintf("%.0f", usuarioID) // Convertir a string

	// Primera consulta: Obtener FkRolesUsuario desde Roles_Usuario**
	rolesUsuarioResponse, err := services.Metodo_get("API_CRUD", "/v1/Roles_Usuario?query=FkUsuarioRoles.Id:", usuarioIDStr)
	if err != nil {
		handleError(c, "Error al obtener el rol del usuario", err)
		return
	}
	fmt.Println("Este es el response de la API Roles_Usuario", string(rolesUsuarioResponse))

	var fkRolesUsuario string
	var rolesUsuarioData map[string]interface{}
	if err := json.Unmarshal(rolesUsuarioResponse, &rolesUsuarioData); err == nil {
		if roles, ok := rolesUsuarioData["Data"].([]interface{}); ok && len(roles) > 0 {
			if role, valid := roles[0].(map[string]interface{}); valid {
				if fkRolObj, exists := role["FkRolesUsuario"].(map[string]interface{}); exists {
					if fkRolID, ok := fkRolObj["Id"].(float64); ok {
						fkRolesUsuario = fmt.Sprintf("%.0f", fkRolID)
					}
				}
			}
		}
	}

	fmt.Println("este es el fkRolesUsuario: ", fkRolesUsuario)

	if fkRolesUsuario == "" {
		handleError(c, "No se encontró una relación de rol para el usuario", nil)
		return
	}
	fmt.Println("FkRolesUsuario obtenido:", fkRolesUsuario)

	// Segunda consulta: Obtener Nombre del rol desde Roles
	rolesResponse, err := services.Metodo_get("API_CRUD", "/v1/Roles?query=Id:", fkRolesUsuario)
	if err != nil {
		handleError(c, "Error al obtener el nombre del rol", err)
		return
	}

	var rolUsuario string
	var rolesData map[string]interface{}
	if err := json.Unmarshal(rolesResponse, &rolesData); err == nil {
		if roles, ok := rolesData["Data"].([]interface{}); ok && len(roles) > 0 {
			if role, valid := roles[0].(map[string]interface{}); valid {
				if nombreRol, exists := role["Nombre"].(string); exists {
					rolUsuario = nombreRol
				}
			}
		}
	}

	if rolUsuario == "" {
		handleError(c, "No se encontró el nombre del rol asociado", nil)
		return
	}
	fmt.Println("Rol del usuario:", rolUsuario)
	// Si la validación es correcta, generar un token JWT
	token, err := jwt.GenerateJWT(int(usuarioID), correo, rolUsuario) //  generar un token JWT basado en los datos del usuario
	if err != nil {
		handleError(c, "Error al generar token", err)
		return
	}
	fmt.Println(token)

	fmt.Println("Login exitoso para el usuario:", correo)
	c.Data["json"] = map[string]string{
		"mensaje": "Login exitoso",
		"token":   token,
	}
	c.ServeJSON()
}

// @Title ValidarToken
// @Description Verifica si el token es válido
// @Param	body	body	map[string]string	true	"Token y correo electrónico"
// @Success 200 {object} map[string]string "Token válido"
// @Failure 400 "Token inválido"
// @Failure 403 "Faltan parámetros"
// @router /validartoken [post]
func (c *AuthController) ValidarToken() {
	fmt.Println("Función POST: Validar Token")

	var datos map[string]string
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &datos); err != nil {
		handleError(c, "Error al leer el cuerpo de la solicitud", err)
		return
	}

	tokenIngresado := datos["token"]
	correo := datos["correo"]
	fmt.Println("token ingresado: ", tokenIngresado, "correo ingresado: ", correo)

	if tokenIngresado == "" || correo == "" {
		handleError(c, "Debe proporcionar el token y el correo electrónico", nil)
		return
	}

	usuario, err := obtenerUsuarioPorCorreo(correo)
	if err != nil {
		handleError(c, "El usuario no existe", err)
		return
	}

	credencialID, err := obtenerFkCredencialID(usuario)
	if err != nil {
		handleError(c, "El usuario no tiene credenciales asociadas", err)
		return
	}
	fmt.Println("ID de la credencial: ", credencialID)

	credencial, err := obtenerCredencialPorID(credencialID)
	if err != nil {
		handleError(c, "Error al obtener la credencial del usuario", err)
		return
	}
	fmt.Println("Esta es la credencial actual: ", credencial)

	tokenGuardado, tieneTokenGuardado := credencial["Token"].(string)
	if !tieneTokenGuardado {
		handleError(c, "No hay token almacenado para este usuario", nil)
		return
	}

	if !services.VerificarToken(tokenIngresado, tokenGuardado) {
		handleError(c, "Token inválido o expirado", nil)
		return
	}
	fmt.Println("Token ingresado:", tokenIngresado)
	fmt.Println("Token guardado:", tokenGuardado)

	fmt.Println("Token válido, procediendo con el cambio de contraseña...")

	c.Data["json"] = map[string]string{"mensaje": "Token válido"}
	c.ServeJSON()
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

// GeneraryEnviarToken
// @Title GeneraryEnviarToken
// @Description update the Usuario_controller.Go
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Usuario_controller.Go	true		"body for Usuario_controller.Go content"
// @Success 200 {object} models.Usuario_controller.Go
// @Failure 403 :id is not int
// @router /:correo [put]
func (c *AuthController) GeneraryEnviarToken() {
	fmt.Println("Función PUT")
	var body map[string]interface{}

	// Leer el cuerpo de la solicitud
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &body); err != nil {
		handleError(c, "Solicitud inválida", err)
		return
	}
	fmt.Println("Este es el body de ingreso:", body)

	// Validar y obtener correo
	correo, ok := body["CorreoElectronico"].(string)
	if !ok {
		handleError(c, "Campo CorreoElectronico inválido", nil)
		return
	}

	// Buscar usuario por correo
	usuario, err := obtenerUsuarioPorCorreo(correo)
	if err != nil {
		handleError(c, "El usuario no existe", err)
		return
	}
	fmt.Println("Usuario encontrado:", usuario)

	// Obtener ID de credencial
	credencialID, err := obtenerFkCredencialID(usuario)
	if err != nil {
		handleError(c, "El usuario no tiene credenciales asociadas", err)
		return
	}
	fmt.Println("ID de la credencial:", credencialID)

	// Generar y almacenar token
	token, hashedToken, err := services.GenerarToken()
	if err != nil {
		handleError(c, "Error al generar el token de recuperación", err)
		return
	}
	fmt.Println("Token generado:", token)

	// Actualizar la credencial con el token
	if err := actualizarTokenCredencial(credencialID, hashedToken); err != nil {
		handleError(c, "Error al guardar el token de recuperación", err)
		return
	}

	// Enviar correo con el token
	if err := services.EnviarCorreo(correo, token); err != nil {
		handleError(c, "Error al enviar el correo de recuperación", err)
		return
	}

	// Enviar respuesta exitosa
	c.Data["json"] = map[string]string{"mensaje": "Correo de recuperación enviado correctamente"}
	c.ServeJSON()

}

// ActualizarContraseña
// @Title ActualizarContraseña
// @Description update the Usuario_controller.Go
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Usuario_controller.Go	true		"body for Usuario_controller.Go content"
// @Success 200 {object} models.Usuario_controller.Go
// @Failure 403 :id is not int
// @router /recuperar/:token [put]
func (c *AuthController) ActualizarContraseña() {
	fmt.Println("Función PUT: Actualizar Contraseña")
	var body map[string]interface{}
	var responseCredencial []byte

	// Leer el cuerpo de la solicitud
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &body); err != nil {
		handleError(c, "Solicitud inválida", err)
		return
	}
	fmt.Println("Este es el body de ingreso:", body)

	// Validar que la nueva contraseña esté en la solicitud
	nuevaContraseña, tieneNuevaContraseña := body["contraseña"].(string)
	if !tieneNuevaContraseña {
		handleError(c, "Debe proporcionar la nueva contraseña", nil)
		return
	}
	fmt.Println("Nueva contraseña:", nuevaContraseña)

	// Validar correo
	correo, ok := body["correo"].(string)
	if !ok {
		handleError(c, "Campo 'correo' inválido o no proporcionado", nil)
		return
	}

	// Buscar usuario por correo
	usuario, err := obtenerUsuarioPorCorreo(correo)
	if err != nil {
		handleError(c, "El usuario no existe", err)
		return
	}
	fmt.Println("Usuario encontrado:", usuario)

	// Obtener ID de credencial actual
	credencialID, err := obtenerFkCredencialID(usuario)
	if err != nil {
		handleError(c, "El usuario no tiene credenciales asociadas", err)
		return
	}
	fmt.Println("ID de la credencial actual:", credencialID)

	// Desactivar la credencial actual
	if err := desactivarCredencial(credencialID); err != nil {
		handleError(c, "Error al desactivar la contraseña actual", err)
		return
	}
	fmt.Println("Credencial actual desactivada")

	// Hashear la contraseña
	hashedNewPass, err := services.HashContraseña(nuevaContraseña)
	if err != nil {
		fmt.Println(err)
		c.Data["json"] = map[string]interface{}{"error": "Error interno al procesar la contraseña"}
		c.ServeJSON()
		return
	}

	// Crear nueva credencial con la nueva contraseña
	jsonCredencial := map[string]interface{}{
		"contraseña": hashedNewPass,
	}
	fmt.Println("JSON para credenciales:", jsonCredencial)

	jsonCredencialByte, _ := json.Marshal(jsonCredencial)
	responseCredencial, err = services.Metodo_post("API_CRUD", "/v1/Credenciales", jsonCredencialByte)
	if err != nil {
		handleError(c, "Error al crear la nueva credencial", err)
		return
	}
	fmt.Println("Respuesta de la API (Credenciales):", string(responseCredencial))

	// Obtener el ID de la nueva credencial
	var credencialResponse map[string]interface{}
	if err := json.Unmarshal(responseCredencial, &credencialResponse); err != nil {
		handleError(c, "Error al parsear la respuesta de credenciales", err)
		return
	}

	// Extraer el ID de la nueva credencial
	newCredencialID, ok := credencialResponse["Data"].(map[string]interface{})["Id"].(float64)
	if !ok {
		handleError(c, "No se encontró el ID en la respuesta de Credenciales", nil)
		return
	}
	fmt.Println("Nuevo ID de credencial:", newCredencialID)

	// Asociar nueva credencial al usuario
	if err := actualizarFkCredencialUsuario(usuario["Id"], newCredencialID); err != nil {
		handleError(c, "Error al actualizar la credencial del usuario", err)
		return
	}

	// Respuesta de éxito
	fmt.Println("Contraseña actualizada con éxito")
	c.Data["json"] = map[string]string{"mensaje": "Contraseña actualizada con éxito"}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the Usuario_controller.Go
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *AuthController) Delete() {
	// Obtener el ID del usuario de la URL
	idUsuario := c.Ctx.Input.Param(":id")
	fmt.Println("ID de usuario a eliminar:", idUsuario)

	// Verificar si el usuario existe en la base de datos
	responseUsuario, err := services.Metodo_get("API_CRUD", "/v1/Usuario/", idUsuario)
	if err != nil || len(responseUsuario) == 0 {
		fmt.Println("Error: Usuario no encontrado o fallo en la consulta")
		c.Data["json"] = map[string]interface{}{
			"error": "Usuario no encontrado",
		}
		c.ServeJSON()
		return
	}

	// Extraer información del usuario
	var usuario map[string]interface{}
	if err := json.Unmarshal(responseUsuario, &usuario); err != nil {
		fmt.Println("Error al parsear respuesta de usuario:", err)
		return
	}

	// Extraer el ID de la credencial asociada
	var credencialID float64
	if data, ok := usuario["Data"].(map[string]interface{}); ok {
		if fkCredencial, exists := data["FkCredencial"].(map[string]interface{}); exists {
			credencialID = fkCredencial["Id"].(float64)
		}
	} else {
		fmt.Println("Error: No se encontró la credencial asociada")
	}

	// Obtener los roles del usuario
	rolesResponse, err := services.Metodo_get("API_CRUD", "/v1/Roles_Usuario?query=FkUsuarioRoles.Id:", idUsuario)
	if err != nil {
		fmt.Println("Error al obtener roles del usuario:", err)
	} else {
		// Extraer IDs de los roles y eliminarlos uno por uno
		var rolesResponseData map[string]interface{}
		if err := json.Unmarshal(rolesResponse, &rolesResponseData); err == nil {
			if rolesData, ok := rolesResponseData["Data"].([]interface{}); ok {
				for _, r := range rolesData {
					if role, valid := r.(map[string]interface{}); valid {
						if roleID, exists := role["Id"].(float64); exists {
							roleIDStr := fmt.Sprintf("%.0f", roleID)
							_, err = services.Metodo_delete("API_CRUD", "/v1/Roles_Usuario/", roleIDStr)
							if err != nil {
								fmt.Println("Error al eliminar el rol:", roleIDStr, err)
							}
						}
					}
				}
			} else {
				fmt.Println("Error: Formato de respuesta inesperado en roles.")
			}
		} else {
			fmt.Println("Error al parsear la respuesta de roles:", err)
		}
	}

	//Eliminar el usuario de la tabla Usuario
	_, err = services.Metodo_delete("API_CRUD", "/v1/Usuario/", idUsuario)
	if err != nil {
		fmt.Println("Error al eliminar usuario:", err)
		c.Data["json"] = map[string]interface{}{
			"error": "No se pudo eliminar el usuario",
		}
		c.ServeJSON()
		return
	}

	// Eliminar la credencial asociada
	if credencialID > 0 {
		_, err = services.Metodo_delete("API_CRUD", "/v1/Credenciales/", fmt.Sprintf("%.0f", credencialID))
		if err != nil {
			fmt.Println("Error al eliminar credenciales del usuario:", err)
		}
	}
	// **Respuesta exitosa**
	c.Data["json"] = map[string]interface{}{"message": "Usuario eliminado exitosamente"}
	c.ServeJSON()
}


// Funciones auxiliares
// Funciones auxiliares para obtener datos de usuario y credenciales
func obtenerUsuarioPorCorreo(correo string) (map[string]interface{}, error) {
	queryParam := "?query=Email:" + correo
	responseUsuario, err := services.Metodo_get("CRUD_SPY", "Usuarios", queryParam)
	if err != nil {
		return nil, err
	}

	var resultado map[string]interface{}
	if err := json.Unmarshal(responseUsuario, &resultado); err != nil {
		return nil, err
	}

	data, ok := resultado["Data"].([]interface{})
	if !ok || len(data) == 0 {
		return nil, errors.New("usuario no encontrado")
	}
	return data[0].(map[string]interface{}), nil
}

// Función para obtener el ID de la credencial asociada al usuario
// Esta función asume que el usuario tiene una relación con la credencial
func obtenerFkCredencialID(Usuarios map[string]interface{}) (string, error) {
	fkCredencial, ok := Usuarios["IdContrasenafk"].(map[string]interface{})
	if !ok {
		return "", errors.New("no se encontró FkCredencial")
	}

	credencialID, ok := fkCredencial["Id"].(float64)
	if !ok {
		return "", errors.New("id de credencial inválido")
	}
	credencialIDInt := int(credencialID)
	return strconv.Itoa(credencialIDInt), nil

}

// Función para actualizar el token de la credencial
// Esta función asume que el ID de la credencial es válido y existe en la base de datos
func actualizarTokenCredencial(credencialID string, hashedToken string) error {
	updateData := map[string]interface{}{"token": hashedToken}
	updateJSON, _ := json.Marshal(updateData)
	_, err := services.Metodo_patch("CRUD_SPY", "Credenciales", credencialID, updateJSON)
	return err
}

// Función para obtener la credencial por su ID
func obtenerCredencialPorID(credencialID string) (map[string]interface{}, error) {
	response, err := services.Metodo_get("CRUD_SPY", "Credenciales/"+credencialID, "")
	if err != nil {
		return nil, err
	}

	var resultado map[string]interface{}
	if err := json.Unmarshal(response, &resultado); err != nil {
		return nil, err
	}

	data, ok := resultado["Data"].(map[string]interface{})
	if !ok {
		return nil, errors.New("credencial no encontrada")
	}

	return data, nil
}

// Función para desactivar la credencial
// Esta función asume que el ID de la credencial es válido y existe en la base de datos
func desactivarCredencial(credencialID string) error {
	updateData := map[string]interface{}{"activo": false, "token": ""}
	updateJSON, _ := json.Marshal(updateData)
	campoactivo, err := services.Metodo_patch("CRUD_SPY", "Credenciales", credencialID, updateJSON)
	fmt.Println("Este es el reponse de la actualizacion del campo activo:", string(campoactivo))
	return err
}

// Función para actualizar la credencial del usuario
// Esta función asume que el ID del usuario y el nuevo ID de credencial son válidos
func actualizarFkCredencialUsuario(usuarioID, newCredencialID interface{}) error {
	updateUserData := map[string]interface{}{
		"FkCredencial": map[string]interface{}{"Id": newCredencialID},
	}
	updateUserJSON, _ := json.Marshal(updateUserData)
	fmt.Println("JSON enviado al PATCH de usuario:", string(updateUserJSON))

	actualizaruser, err := services.Metodo_patch("API_CRUD", "/v1/Usuario", fmt.Sprintf("%v", usuarioID), updateUserJSON)
	fmt.Println("este es el respone de actulizar id de usuario", string(actualizaruser))
	return err

}

// Función para manejar errores y enviar respuestas JSON
// Esta función se encarga de establecer el estado de la respuesta y enviar un mensaje de error en formato JSON
func handleError(c *AuthController, mensaje string, err error) {
	c.Ctx.Output.SetStatus(400)
	c.Data["json"] = map[string]string{"error": mensaje}
	c.ServeJSON()
}