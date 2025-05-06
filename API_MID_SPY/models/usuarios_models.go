package models

import (
	
)


type UserRequest struct {
	Nombres              string `json:"Nombres"`
	Apellidos            string `json:"Apellidos"`
	NumeroIdentificacion string `json:"NumeroIdentificacion"`
	Email                string `json:"Email"`
	Telefono             int    `json:"Telefono"`
	Direccion            string `json:"Direccion"`
	IdRolesFk            struct {
		Id     int    `json:"Id"`
		Roles string `json:"Roles"`
		Cargos string `json:"Cargos"`
	} `json:"IdRolesFk"`
	IdContrasenaFk struct {
		Contrasena string `json:"Contrasena"`
	} `json:"IdContrasenaFk"`
}

type Alert struct {
	Type string
	Code string
	Body interface{}
}