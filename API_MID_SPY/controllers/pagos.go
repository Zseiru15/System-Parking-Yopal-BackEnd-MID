package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/astaxie/beego"
	"github.com/sena_2824182/System-Parking-Yopal-BackEnd-MID/API_MID_SPY/services"
)

// PagosController operations for Pagos
type PagosController struct {
	beego.Controller
}

// URLMapping ...
func (c *PagosController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Create
// @Description create Pagos
// @Param	body		body 	models.Pagos	true		"body for Pagos content"
// @Success 201 {object} models.Pagos
// @Failure 403 body is empty
// @router / [post]
func (c *PagosController) Post() {
	fmt.Println("Post Vacio")
}

// GetOne ...
// @Title GetOne
// @Description get Pagos by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Pagos
// @Failure 403 :id is empty
// @router /:id [get]
func (c *PagosController) GetOne() {
	fmt.Println("GetOne Vacio")
}

// GetAll ...
// @Title GetAll
// @Description get Pagos
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Pagos
// @Failure 403
// @router / [get]
func (c *PagosController) GetAll() {
	body, err := services.Metodo_get("CRUD_SPY", "Pagos", "?limit=0")
	if err != nil || len(body) == 0 {
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  500,
			"Message": "Error al obtener los pagos",
			"error":   err.Error(),
		}
		c.ServeJSON()
		return
	}

	var responseData map[string]interface{}
	if err := json.Unmarshal(body, &responseData); err != nil {
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  500,
			"Message": "Error al procesar la respuesta del servidor",
			"error":   err.Error(),
		}
		c.ServeJSON()
		return
	}

	paymentsArray, ok := responseData["data"].([]interface{})
	if !ok {
		c.Data["json"] = map[string]interface{}{
			"Success": false,
			"Status":  500,
			"Message": "Estructura de datos incorrecta",
		}
		c.ServeJSON()
		return
	}

	var pagos []map[string]interface{}
	for _, item := range paymentsArray {
		payments, ok := item.(map[string]interface{})
		if !ok {
			continue
		}

		// Agregar usuario procesado a la lista final
		pagos = append(pagos, map[string]interface{}{
			"Id":            payments["Id"],
			"PayPalOrdenID": payments["PayPalOrderID"],
			"Amount":        payments["Amount"],
			"Currency":      payments["Currency"],
			"Status":        payments["Status"],
			"PayerEmail":    payments["PayerEmail"],
			"ReceiverEmail": payments["ReceiverEmail"],
			"CreatedAt":     payments["CreatedAt"],
		})
	}

	// Respuesta JSON optimizada
	c.Data["json"] = map[string]interface{}{
		"Success": true,
		"Status":  200,
		"Message": "Consulta exitosa",
		"Data":    pagos,
		"Total":   len(pagos), // Indica cuántos usuarios se obtuvieron
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the Pagos
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Pagos	true		"body for Pagos content"
// @Success 200 {object} models.Pagos
// @Failure 403 :id is not int
// @router /:id [put]
func (c *PagosController) Put() {
	fmt.Println("Put Vacio")
}

// Delete ...
// @Title Delete
// @Description delete the Pagos
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *PagosController) Delete() {
	fmt.Println("Delete Vacio")
}


func GetPayPalAccessToken() (string, error) {
	clientID := os.Getenv("AQvAF1VdHXee7eTPVYy515ni8mkeB698nAP3vom3ZBuBLTHEaci2e9ySj1_HMFniutieNbi-8Y0RRxCi")
	secret := os.Getenv("EE-O2r2HMVK_SpLej1R5xFjZ3lkVMfiaeIkip3m01PWe3Pa57dyHFUhycDMEdz1L4Kd9y0Eca7UkcpXV")
	url := "https://api-m.sandbox.paypal.com/v1/oauth2/token" // Cambia a live para producción

	reqBody := strings.NewReader("grant_type=client_credentials")
	req, err := http.NewRequest("POST", url, reqBody)
	if err != nil {
		return "", err
	}

	req.SetBasicAuth(clientID, secret)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return "", errors.New("No se pudo obtener token de PayPal")
	}

	body, _ := ioutil.ReadAll(resp.Body)

	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return "", err
	}

	accessToken, ok := result["access_token"].(string)
	if !ok {
		return "", errors.New("Token inválido")
	}

	return accessToken, nil
}
