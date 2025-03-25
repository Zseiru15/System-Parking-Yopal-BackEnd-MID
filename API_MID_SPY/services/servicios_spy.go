package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/astaxie/beego"
)

// Mostrar el valor de resultado 1
func ProcesarJsonArreglos(datos []byte) ([]map[string]interface{}, error) {
	var result []map[string]interface{}

	//Intentamos deserializar el []byte en un []map[]map[string]interface{}
	err := json.Unmarshal(datos, &result)
	if err != nil {
		return nil, err
	}
	//Retornamos el slice de mapas
	return result, nil
}

// arreglar el orden de parametros seleccionados
func Metodo_get(nombre_servicio, endpoint, parametro string) ([]byte, error) {
	url := beego.AppConfig.String(nombre_servicio)+ endpoint +"/"+ parametro
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return body, nil
}

func Metodo_post(nombre_servicio string, data []byte) ([]byte, error) {
	url := beego.AppConfig.String(nombre_servicio)
	response, err := http.Post(url, "application/json", bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	return body, nil
}

func ProcesarJson(datos []byte) (map[string]interface{}, error) {
	var result map[string]interface{}
	//Intentamos deserializar el []byte en un []map[]map[string]interface{}
	err2 := json.Unmarshal(datos, &result)
	if err2 != nil {
		log.Fatal(err2)
		return nil, err2
	}
	//Retornamos el slice de mapas
	return result, nil
}

func Metodo_put(nombre_servicio string, id string, data []byte) ([]byte, error) {
	//Obtener la URL base desde la configuracion de Beego
	baseURL := beego.AppConfig.String(nombre_servicio)

	//Construir la URL final con ID
	url := fmt.Sprintf("%s/%s", baseURL, id)

	//Crear la solicitud PUT
	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	//Establecer el encabezado Content-Type
	req.Header.Set("Content-Type", "application/json")

	//Enviar la solicitud
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	//Leer la respuesta
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	return body, nil
}