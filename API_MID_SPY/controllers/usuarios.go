package controllers

import (
	"github.com/astaxie/beego"
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

	//asignacion de datos al body
	body, _ := services.Metodo_get("Servicio_Cartas", id_ingreso)
	body2, _ := services.Metodo_get("Servicio_Usuarios", id_ingreso)

	resultado1, _ := services.ProcesarJsonArreglos(body)
	//----------------------------------------------------------------------------------------

	//var result map[string]interface{}  // El JSON que esperas es un array de objetos
	var result2 map[string]interface{} // El JSON que esperas es un array de objetos

	//err = json.Unmarshal(body, &result)
	//if err != nil {
	//	log.Fatal("Error al parsear JSON:", err)
	//}

	err2 := json.Unmarshal(body2, &result2)
	if err2 != nil {
		log.Fatal("Error al parsear JSON:", err2)
	}

	//agregar un campo nuevo
	for i := range resultado1 {
		resultado1[i] = map[string]interface{}{
			"Campo_nuevo": i + 1,
			"body":        resultado1[i]["body"],
		}
	}

	result2 = map[string]interface{}{
		"direccion":          result2["address"],
		"telefono":           result2["phone"],
		"codigo_postal":      result2["address"].(map[string]interface{})["zipcode"],
		"dirrecion_telefono": map[string]interface{}{"dirrecion": result2["address"], "telefono": result2["phone"]},
	}

	//------------------------------------------------

	//Sacar una parte de un json del resultado 1
	resultado := append(resultado1, result2)

	//informacion de estado
	//fmt.Println("La cantidad de datos son", len(resultado))
	c.Data["json"] = map[string]interface{}{
		"Succes":          true,
		"Status":          200,
		"Message":         "Consulta existosa",
		"Data":            resultado,
		"Cantidad Cartas": len(resultado)}
	c.ServeJSON()
}

// GetAll ...
// @Title GetAll
// @Description get Usuarios
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Usuarios
// @Failure 403
// @router / [get]
func (c *UsuariosController) GetAll() {

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
