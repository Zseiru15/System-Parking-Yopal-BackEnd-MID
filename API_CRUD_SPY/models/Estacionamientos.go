package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type Estacionamientos struct {
	Id                  int                       `orm:"column(Id_estacionamientos);pk"`
	Nombres             string                    `orm:"column(Nombres)"`
	Direccion           string                    `orm:"column(Direccion)"`
	Email               string                    `orm:"column(Email)"`
	Telefono            float64                   `orm:"column(Telefono)"`
	Capacidad           float64                   `orm:"column(Capacidad)"`
	Estado              bool                      `orm:"column(Estado)"`
	IdAdministradoresFk *AdministradoresEmpleados `orm:"column(Id_Administradores_fk);rel(fk)"`
	Largo               string                    `orm:"column(Largo)"`
	Ancho               string                    `orm:"column(Ancho)"`
	Altura              string                    `orm:"column(Altura)"`
	Descripcion         string                    `orm:"column(Descripcion)"`
	FechaRegistro       time.Time                 `orm:"column(Fecha_Registro);type(timestamp with time zone)"`
	FechaInicio         time.Time                 `orm:"column(Fecha_Inicio);type(timestamp with time zone)"`
	FechaModificacion   time.Time                 `orm:"column(Fecha_Modificacion);type(timestamp with time zone)"`
	FechaFinal          time.Time                 `orm:"column(Fecha_Final);type(timestamp with time zone)"`
	Latitud             float64                   `orm:"column(Latitud)"`
	Longitud            float64                   `orm:"column(Longitud)"`
}

func (t *Estacionamientos) TableName() string {
	return "Estacionamientos"
}

func init() {
	orm.RegisterModel(new(Estacionamientos))
}

// AddEstacionamientos insert a new Estacionamientos into database and returns
// last inserted Id on success.
func AddEstacionamientos(m *Estacionamientos) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetEstacionamientosById retrieves Estacionamientos by Id. Returns error if
// Id doesn't exist
func GetEstacionamientosById(id int) (v *Estacionamientos, err error) {
	o := orm.NewOrm()
	v = &Estacionamientos{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllEstacionamientos retrieves all Estacionamientos matches certain condition. Returns empty list if
// no records exist
func GetAllEstacionamientos(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Estacionamientos))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []Estacionamientos
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateEstacionamientos updates Estacionamientos by Id and returns error if
// the record to be updated doesn't exist
func UpdateEstacionamientosById(m *Estacionamientos) (err error) {
	o := orm.NewOrm()
	v := Estacionamientos{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteEstacionamientos deletes Estacionamientos by Id and returns error if
// the record to be deleted doesn't exist
func DeleteEstacionamientos(id int) (err error) {
	o := orm.NewOrm()
	v := Estacionamientos{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Estacionamientos{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
