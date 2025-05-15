package controllers

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
	"github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_MID_SPY/security"
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

func (c *UsuariosController) Login() {
	var body map[string]string

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &body); err != nil {
		c.CustomAbort(400, "Cuerpo inválido: "+err.Error())
		return
	}

	email := body["email"]
	password := body["password"]

	if email == "" || password == "" {
		c.CustomAbort(400, "Email y contraseña son requeridos")
		return
	}

	// Consulta al CRUD
	resp, err := services.Metodo_get("CRUD_SPY", "usuarios/login/"+email, "")
	if err != nil || len(resp) == 0 {
		c.CustomAbort(500, "Error al obtener el usuario o el usuario no existe")
		return
	}

	var respMap map[string]interface{}
	if err := json.Unmarshal(resp, &respMap); err != nil {
		c.CustomAbort(500, "Error procesando la respuesta del CRUD")
		return
	}

	// Verificación segura de campos anidados
	data, ok := respMap["data"].(map[string]interface{})
	if !ok {
		c.CustomAbort(500, "Respuesta inválida del CRUD: data no existe")
		return
	}

	credenciales, ok := data["IdContrasenaFk"].(map[string]interface{})
	if !ok {
		c.CustomAbort(500, "Credenciales no encontradas")
		return
	}

	passwordStoredRaw, ok := credenciales["Contrasena"]
	if !ok {
		c.CustomAbort(500, "Campo contraseña no disponible")
		return
	}

	passwordStored := strings.TrimSpace(fmt.Sprintf("%v", passwordStoredRaw))

	// Comparación directa (sin encriptación)
	if password != passwordStored {
		c.CustomAbort(401, "Contraseña incorrecta")
		return
	}

	// Eliminar campo sensible antes de responder
	delete(data, "IdContrasenaFk")

	c.Data["json"] = map[string]interface{}{
		"Success": true,
		"Status":  200,
		"Message": "Login exitoso",
		"Data":    data,
	}
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
		"Id":                   userData["data"].(map[string]interface{})["Id"],
		"Nombres":              userData["data"].(map[string]interface{})["Nombres"],
		"Apellidos":            userData["data"].(map[string]interface{})["Apellidos"],
		"NumeroIdentificacion": userData["data"].(map[string]interface{})["NumeroIdentificacionUsuarios"],
		"Email":                userData["data"].(map[string]interface{})["Email"],
		"Telefono":             userData["data"].(map[string]interface{})["Telefono"],
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
			"Id":                   user["Id"],
			"Nombres":              user["Nombres"],
			"Apellidos":            user["Apellidos"],
			"NumeroIdentificacion": user["NumeroIdentificacionUsuarios"],
			"Email":                user["Email"],
			"Telefono":             user["Telefono"],
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
// @Title Update
// @Description update Usuarios
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body	body 	models.Usuarios	true		"body for Usuarios content"
// @Success 200 {object} models.Usuarios
// @Failure 400 Invalid ID
// @Failure 500 Error updating user
// @router /:id [put]
func (c *UsuariosController) Put() {
	id_ingreso := c.Ctx.Input.Param(":id") // Captura el ID desde la URL

	// Decodificar el cuerpo de la solicitud
	var body_actualizacion map[string]interface{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &body_actualizacion); err != nil {
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  400,
			"Message": "Error al procesar el cuerpo de la solicitud",
		}
		c.ServeJSON()
		return
	}

	// Convertir body_actualizacion a JSON antes de enviarlo a Metodo_put
	json_usuario_byte, err := json.Marshal(body_actualizacion)
	if err != nil {
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  500,
			"Message": "Error al convertir datos a JSON",
		}
		c.ServeJSON()
		return
	}

	// Llamar a Metodo_put para actualizar el usuario en CRUD_SPY
	response_usuario, err := services.Metodo_put("CRUD_SPY", "Usuarios", id_ingreso, json_usuario_byte)
	if err != nil {
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  500,
			"Message": "Error al actualizar el usuario",
		}
		c.ServeJSON()
		return
	}

	// Decodificar la respuesta del servicio
	var temporal_usuario map[string]interface{}
	if err := json.Unmarshal(response_usuario, &temporal_usuario); err != nil {
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  500,
			"Message": "Error al procesar la respuesta del servidor",
		}
		c.ServeJSON()
		return
	}

	// 📌 IMPRIMIR LA RESPUESTA COMPLETA
	fmt.Println("Respuesta del servidor CRUD_SPY:", temporal_usuario)

	// Validar si "data" está presente en la respuesta
	data, ok := temporal_usuario["data"].(map[string]interface{})
	if !ok {
		c.Data["json"] = map[string]interface{}{
			"Success":     false,
			"Status":      500,
			"Message":     "Estructura de datos incorrecta en la respuesta del servidor",
			"RawResponse": temporal_usuario, // 🔍 Incluir la respuesta completa para depuración
		}
		c.ServeJSON()
		return
	}

	// Respuesta JSON optimizada
	c.Data["json"] = map[string]interface{}{
		"Succes":  true,
		"Status":  200,
		"type":    "put",
		"Message": "Actualización exitosa",
		"Data":    data,
	}
	c.ServeJSON()
}

// Delete ...
// @Title Disable
// @Description Cambia el estado de Usuarios a false en lugar de eliminarlo
// @Param	id		path 	string	true		"The ID of the user to disable"
// @Success 200 {string} Usuario deshabilitado exitosamente
// @Failure 400 ID no válido
// @Failure 404 Usuario no encontrado
// @Failure 500 Error interno del servidor
// @router /:id [delete]
func (c *UsuariosController) Delete() {
	id_ingreso := c.Ctx.Input.Param(":id") // Capturar ID de la URL

	if id_ingreso == "" {
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  400,
			"Message": "ID no proporcionado",
		}
		c.ServeJSON()
		return
	}

	// 📌 Obtener el usuario antes de modificarlo
	get_inicial, err := services.Metodo_get("CRUD_SPY", "Usuarios", id_ingreso)
	if err != nil || len(get_inicial) == 0 {
		fmt.Println("Usuario no encontrado en el GET inicial.")
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  404,
			"Message": "Usuario no encontrado",
		}
		c.ServeJSON()
		return
	}

	// Decodificar JSON recibido
	var usuarioActual map[string]interface{}
	if err := json.Unmarshal(get_inicial, &usuarioActual); err != nil {
		fmt.Println("Error al procesar el usuario:", err)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  500,
			"Message": "Error al procesar el usuario",
		}
		c.ServeJSON()
		return
	}

	// 📌 Verificar que el usuario tiene la clave "Estado"
	if _, ok := usuarioActual["data"].(map[string]interface{})["Estado"]; !ok {
		fmt.Println("El usuario no tiene un campo 'Estado'.")
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  500,
			"Message": "El usuario no tiene un campo 'Estado'.",
		}
		c.ServeJSON()
		return
	}

	// 📌 Crear JSON con la actualización del estado
	json_nuevo := map[string]interface{}{
		"Estado": false, // Solo cambiamos el estado a false
	}

	// Convertir a JSON
	json_byte, _ := json.Marshal(json_nuevo)

	// 📌 Imprimir JSON antes de enviarlo
	fmt.Println("📤 JSON a enviar en PUT:", string(json_byte))

	// Llamar a Metodo_put para actualizar el estado
	response_put, err := services.Metodo_put("CRUD_SPY", "Usuarios", id_ingreso, json_byte)
	if err != nil {
		fmt.Println("Error en Metodo_put:", err)
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  500,
			"Message": "Error al deshabilitar el usuario",
		}
		c.ServeJSON()
		return
	}

	// 📌 Ver respuesta del CRUD_SPY
	fmt.Println("Respuesta del CRUD_SPY al PUT:", string(response_put))

	// Confirmar la actualización
	c.Data["json"] = map[string]interface{}{
		"Success": true,
		"Status":  200,
		"Message": "Usuario deshabilitado correctamente",
	}
	c.ServeJSON()
}
