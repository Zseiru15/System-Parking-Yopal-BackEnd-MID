package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/astaxie/beego"
)

// Orden de los servicios
// 1. ProcesarJsonArreglos
// 2. ProcesarJson
// 3. Metodo_post
// 4. Metodo_get
// 5. Metodo_put
// 6. Metodo_delete
// 7. Metodo_patch
// 8. GenerarToken
// 9. VerificarToken
// 10. init
// 11. EnviarCorreo
// 12. HashContraseña

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

func ObtenerIDRol(rol string) int {
	id_rol := 0

	switch rol {
	case "Usuario":
		id_rol = 1
	case "Empleado":
		id_rol = 2
	case "Administrador":
		id_rol = 3
	default:
		fmt.Println("Rol no reconocido")
	}

	return id_rol // ✅ correcto porque la función retorna int
}

func ObtenerIDTipo(tipo string) int {
	id_tipo := 0

	switch tipo {
	case "Subterráneo":
		id_tipo = 1
	case "Arire Libre":
		id_tipo = 2
	default:
		fmt.Println("Tipo no reconocido")
	}

	return id_tipo // ✅ correcto porque la función retorna int
}

func ObtenerIDSombras(sombra string) int {
	id_sombra := 0

	switch sombra {
	case "Subterráneo":
		id_sombra = 1
	case "Arire Libre":
		id_sombra = 2
	default:
		fmt.Println("Tipo no reconocido")
	}

	return id_sombra // ✅ correcto porque la función retorna int
}

func Metodo_post(host string, endpoint string, data []byte) ([]byte, error) {

	url := beego.AppConfig.String(host) + endpoint // Construir la URL
	fmt.Println("URL enviada:", url)

	if url == "" {
		return nil, fmt.Errorf("no se encontró la configuración para %s", host)
	}

	// Asegurar que la URL tiene "http://"
	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		url = "http://" + url
	}

	// Enviar la solicitud HTTP POST
	response, err := http.Post(url, "application/json", bytes.NewBuffer(data))

	if err != nil {
		fmt.Println("Error en POST:", err)
		return nil, fmt.Errorf("error en POST a %s: %v", url, err)
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error al leer la respuesta:", err)
		return nil, fmt.Errorf("error al leer la respuesta: %v", err)
		// log.Fatal(err)
	}


	return body, nil
}

// arreglar el orden de parametros seleccionados
func Metodo_get(nombre_servicio, endpoint, parametro string) ([]byte, error) {
	url := beego.AppConfig.String(nombre_servicio) + endpoint + "/" + parametro
	fmt.Println("URL enviada:", url)

	if url == "" {
		return nil, fmt.Errorf("no se encontró la configuración para %s", nombre_servicio)
	}

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error en GET a %s: %v", url, err)
		// return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		return nil, fmt.Errorf("error en API GET: %d - %s", resp.StatusCode, string(body))
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error al leer la respuesta: %v", err)
		// log.Fatal(err)
	}
	return body, nil
}

func Metodo_put(nombre_servicio, endpoint, id string, data []byte) ([]byte, error) {
	// Obtener la URL base desde la configuración de Beego
	baseURL := beego.AppConfig.String(nombre_servicio)
	if baseURL == "" {
		return nil, fmt.Errorf("no se encontró la configuración para %s", nombre_servicio)
	}

	// Asegurar que la URL tiene "http://"
	if !strings.HasPrefix(baseURL, "http://") && !strings.HasPrefix(baseURL, "https://") {
		baseURL = "http://" + baseURL
	}

	// Construir la URL final con el endpoint y el ID
	url := fmt.Sprintf("%s%s/%s", baseURL, endpoint, id)
	fmt.Println("URL enviada:", url)

	// Crear la solicitud PUT
	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("error creando la solicitud PUT: %v", err)
	}
	// Establecer el encabezado Content-Type
	req.Header.Set("Content-Type", "application/json")
	fmt.Println("solicitud PUT: ", req)

	// Enviar la solicitud
	client := &http.Client{}
	response, err := client.Do(req)
	fmt.Println("este es el response al enviar la solicitud", response)
	if err != nil {
		return nil, fmt.Errorf("error al enviar la solicitud PUT: %s: %v", url, err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK && response.StatusCode != http.StatusNoContent {
		body, _ := ioutil.ReadAll(response.Body)
		return nil, fmt.Errorf("error en API PUT: %d - %s", response.StatusCode, string(body))
	}

	// Leer la respuesta
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("error al leer la respuesta del servidor: %v", err)
	}

	// Devolver la respuesta sin usar log.Fatal
	fmt.Println("Respuesta de la API:", string(body))
	return body, nil
}

func Metodo_delete(nombre_servicio, endpoint, parametro string) ([]byte, error) {
	// Construir la URL de eliminación
	baseURL := beego.AppConfig.String(nombre_servicio) + endpoint + "/" + parametro
	if baseURL == "" {
		return nil, fmt.Errorf("no se encontró la configuración para %s", nombre_servicio)
	}

	// Asegurar que la URL tiene "http://"
	if !strings.HasPrefix(baseURL, "http://") && !strings.HasPrefix(baseURL, "https://") {
		baseURL = "http://" + baseURL
	}

	url := fmt.Sprintf("%s%s/%s", baseURL, endpoint, parametro)
	fmt.Println("URL enviada:", url)

	// Crear la solicitud DELETE
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return nil, fmt.Errorf("error al crear la solicitud DELETE: %v", err)
		// return nil, err
	}

	// Enviar la solicitud
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error en DELETE a %s: %v", url, err)
		// return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK && response.StatusCode != http.StatusNoContent {
		body, _ := ioutil.ReadAll(response.Body)
		return nil, fmt.Errorf("error en API DELETE: %d - %s", response.StatusCode, string(body))
	}

	// Leer la respuesta
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("error al leer la respuesta: %v", err)
		// log.Fatal(err)
	}
	fmt.Println("Respuesta de la API:", string(body))
	return body, nil
}

func Metodo_patch(nombre_servicio, endpoint, id string, data []byte) ([]byte, error) {
	baseURL := beego.AppConfig.String(nombre_servicio)
	if baseURL == "" {
		return nil, fmt.Errorf("no se encontró la configuración para %s", nombre_servicio)
	}

	// Asegurar que la URL tiene "http://"
	if !strings.HasPrefix(baseURL, "http://") && !strings.HasPrefix(baseURL, "https://") {
		baseURL = "http://" + baseURL
	}

	url := fmt.Sprintf("%s%s/%s", baseURL, endpoint, id)
	fmt.Println("URL construida:", url)

	// Crear la solicitud PATCH
	req, err := http.NewRequest(http.MethodPatch, url, bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("error al crear la solicitud PATCH: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	fmt.Println("solicitud patch:", req)

	// Enviar la solicitud
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error en PATCH a %s: %v", url, err)
	}
	defer response.Body.Close()
	fmt.Println("ojoo: ", response)

	if response.StatusCode != http.StatusOK && response.StatusCode != http.StatusNoContent {
		body, _ := ioutil.ReadAll(response.Body)
		return nil, fmt.Errorf("error en API PATCH: %d - %s", response.StatusCode, string(body))
	}
	fmt.Println("ojoo: ", response)

	// Leer la respuesta
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("error al leer la respuesta: %v", err)
	}
	fmt.Println("respuesta api", string(body))

	fmt.Println("Respuesta de la API:", string(body))
	return body, nil
}