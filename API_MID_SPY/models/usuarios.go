package models

import (
	"errors"
	"strconv"
	"time"

	"github.com/astaxie/beego/orm"
)

var (
	UserList map[string]*User
)

func init() {
	UserList = make(map[string]*User)
	u := User{"user_11111", "nombre_astaxie", "apellido_astaxie", "11111", Profile{"male", 20, "Singapore", "astaxie@gmail.com", 1}}
	UserList["user_11111"] = &u
}

type User struct {
	Id         string
	Nombres    string `json:"Nombres"`
	Apellidos  string `json:"Apellidos"`
	Contrasena string `json:"Contraseña"`
	Profile    Profile
}

type Profile struct {
	Genero    string `json:"Genero"`
	Edad      int    `json:"Edad"`
	Direccion string `json:"Direccion"`
	Email     string `json:"Email"`
	IdRolesFk int    `json:"IdRolesFk"`
}

func AddUser(u User) string {
	u.Id = "user_" + strconv.FormatInt(time.Now().UnixNano(), 10)
	UserList[u.Id] = &u
	return u.Id
}

func GetUser(uid string) (u *User, err error) {
	if u, ok := UserList[uid]; ok {
		return u, nil
	}
	return nil, errors.New("User not exists")
}

func GetAllUsers() map[string]*User {
	return UserList
}

func UpdateUser(uid string, uu *User) (a *User, err error) {
	if u, ok := UserList[uid]; ok {
		if uu.Nombres != "" {
			u.Nombres = uu.Nombres
		}
		if uu.Apellidos != "" {
			u.Apellidos = uu.Apellidos
		}
		if uu.Contrasena != "" {
			u.Contrasena = uu.Contrasena
		}
		if uu.Profile.Edad != 0 {
			u.Profile.Edad = uu.Profile.Edad
		}
		if uu.Profile.Direccion != "" {
			u.Profile.Direccion = uu.Profile.Direccion
		}
		if uu.Profile.Genero != "" {
			u.Profile.Genero = uu.Profile.Genero
		}
		if uu.Profile.Email != "" {
			u.Profile.Email = uu.Profile.Email
		}
		return u, nil
	}
	return nil, errors.New("User Not Exist")
}

func Login(username, password string) bool {
	for _, u := range UserList {
		if u.Nombres == username && u.Contrasena == password {
			return true
		}
	}
	return false
}

func DeleteUser(uid string) {
	delete(UserList, uid)
}

type UserRequest struct {
	Nombres              string `json:"Nombres"`
	Apellidos            string `json:"Apellidos"`
	NumeroIdentificacion string `json:"NumeroIdentificacion"`
	Edad                 int    `json:"Edad"`
	Email                string `json:"Email"`
	Telefono             int    `json:"Telefono"`
	Direccion            string `json:"Direccion"`
	IdRolesFk            struct {
		Id     int    `json:"Id"`
		Nombre string `json:"Nombre"`
	} `json:"IdRolesFk"`
	IdContrasenaFk struct {
		Id int `json:"Id"`
	} `json:"IdContrasenaFk"`
	FechaNacimiento string `json:"Fecha_nacimiento"` // ✅ Debe ser un string
}

type Alert struct {
	Type string
	Code string
	Body interface{}
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string      `json:"token"`
	User  interface{} `json:"user"`
}

type Usuarios struct {
    IdUsuarios      int       `orm:"column(Id_usuarios);pk;auto"`
    Nombres         string    `orm:"column(Nombres);size(100)"`
    Apellidos       string    `orm:"column(Apellidos);size(100)"`
    NumeroIdentifi  string    `orm:"column(Numero_Identificacion_Usuarios);size(50)"`
    Email           string    `orm:"column(Email);size(100);unique"`
    
    // Relación con Credenciales (clave foránea)
    Credencial      *Credenciales `orm:"column(Id_Contrasena_fk);rel(fk)"`
    
    Telefono        int64     `orm:"column(Telefono)"`
    Direccion       string    `orm:"column(Direccion);size(200)"`
    Edad            int       `orm:"column(Edad)"`
    FechaNacimiento string    `orm:"column(Fecha_nacimiento);size(50)"`
    Estado          bool      `orm:"column(Estado);default(true)"`
    FechaRegistro   time.Time `orm:"column(Fecha_Registro);auto_now_add;type(timestamp)"`
    FechaModifica   time.Time `orm:"column(Fecha_Modifica);auto_now;type(timestamp)"`
    
    // Relación con Roles (clave foránea)
    Rol             *Roles    `orm:"column(Id_Roles_fk);rel(fk)"`
    
    Usuario         string    `orm:"column(Usuario);size(50)"`
    Imagen          string    `orm:"column(Imagen);size(255)"`
}

func (u *Usuarios) TableName() string {
	return "Usuarios"
}

func init() {
    // Registrar todos los modelos
    orm.RegisterModel(
        new(Usuarios),
        new(Credenciales),
        new(Roles),
    )
}
