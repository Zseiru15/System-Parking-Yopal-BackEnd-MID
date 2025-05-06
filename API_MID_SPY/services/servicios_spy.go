package services

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/exp/rand"
	"gopkg.in/gomail.v2"
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

func Metodo_post(nombre_servicio string, endpoint string, data []byte) ([]byte, error) {
	
	url := beego.AppConfig.String(nombre_servicio) + endpoint // Construir la URL
	fmt.Println("URL enviada:", url)

	if url == "" {
		return nil, fmt.Errorf("no se encontró la configuración para %s", nombre_servicio)
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

	fmt.Println("Respuesta de la API:", string(body))

	return body, nil
}

// arreglar el orden de parametros seleccionados
func Metodo_get(nombre_servicio, endpoint, parametro string) ([]byte, error) {
	url := beego.AppConfig.String(nombre_servicio) + endpoint + "/" + parametro
	fmt.Println("URL enviada:", url)

	if url == "" {
		return nil, fmt.Errorf("no se encontró la configuración para %s", nombre_servicio)
	}

	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		url = "http://" + url
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

// GenerarToken crea un token de 5 dígitos aleatorios y lo hashea
func GenerarToken() (string, string, error) {
	token := fmt.Sprintf("%05d", 10000+rand.Intn(90000)) // Token de 5 dígitos

	// Hashear el token antes de guardarlo
	hashedToken, err := bcrypt.GenerateFromPassword([]byte(token), bcrypt.DefaultCost)
	if err != nil {
		return "", "", err
	}

	return token, string(hashedToken), nil
}

// VerificarToken compara el token ingresado con el hash almacenado
func VerificarToken(tokenIngresado string, tokenGuardado string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(tokenGuardado), []byte(tokenIngresado))
	return err == nil
}

func init() {
	err := godotenv.Load() // Carga el archivo .env en el entorno
	if err != nil {
		log.Println("Error cargando el archivo .env:", err)
	}
}

// EnviarCorreo envía un token de recuperación al usuario
func EnviarCorreo(destinatario string, token string) error {
	// Obtener credenciales del .env
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")
	smtpUser := os.Getenv("SMTP_USER")
	smtpPass := os.Getenv("SMTP_PASS")
	fmt.Println("SMTP_PORT:", smtpPort)

	if smtpHost == "" || smtpPort == "" || smtpUser == "" || smtpPass == "" {
		log.Println("Error: Configuración de SMTP incompleta")
		return fmt.Errorf("configuración de SMTP incompleta")
	}

	// Convertir puerto a entero
	port, err := strconv.Atoi(smtpPort)
	if err != nil {
		log.Printf("Error convirtiendo SMTP_PORT a número: %v", err)
		return err
	}
	fmt.Println("SMTP_PORT:", smtpPort)

	// Configurar mensaje
	mensaje := gomail.NewMessage()
	mensaje.SetHeader("From", smtpUser)
	mensaje.SetHeader("To", destinatario)
	mensaje.SetHeader("Subject", "Recuperación de contraseña")
	mensaje.SetBody("text/plain", fmt.Sprintf("Tu código de recuperación es: %s", token))

	// Configurar servidor SMTP
	dialer := gomail.NewDialer(smtpHost, port, smtpUser, smtpPass)
	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true} // Descomentar si hay problemas con TLS

	// Enviar correo
	if err := dialer.DialAndSend(mensaje); err != nil {
		log.Printf("Error enviando el correo: %v", err)
		return err
	}

	fmt.Println("Correo enviado correctamente a", destinatario)
	return nil
}

func HashContraseña(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("error al hashear la contraseña: %v", err)
	}
	return string(hashed), nil
}