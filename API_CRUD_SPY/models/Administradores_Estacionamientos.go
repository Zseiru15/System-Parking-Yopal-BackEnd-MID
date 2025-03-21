package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type AdministradoresEstacionamientos struct {
	Id                   int                       `orm:"column(Id_Administradores_Estacionamientos);pk"`
	IdAdministradoresFk  *AdministradoresEmpleados `orm:"column(Id_Administradores_fk);rel(fk)"`
	IdEstacionamientosFk *Estacionamientos         `orm:"column(Id_Estacionamientos_fk);rel(fk)"`
	FechaRegistro        time.Time                 `orm:"column(Fecha_Registro);type(timestamp with time zone)"`
	FechaModificacion    time.Time                 `orm:"column(Fecha_Modificacion);type(timestamp with time zone)"`
}

func (t *AdministradoresEstacionamientos) TableName() string {
	return "Administradores_Estacionamientos"
}

func init() {
	orm.RegisterModel(new(AdministradoresEstacionamientos))
}

// AddAdministradoresEstacionamientos insert a new AdministradoresEstacionamientos into database and returns
// last inserted Id on success.
func AddAdministradoresEstacionamientos(m *AdministradoresEstacionamientos) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetAdministradoresEstacionamientosById retrieves AdministradoresEstacionamientos by Id. Returns error if
// Id doesn't exist
func GetAdministradoresEstacionamientosById(id int) (v *AdministradoresEstacionamientos, err error) {
	o := orm.NewOrm()
	v = &AdministradoresEstacionamientos{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllAdministradoresEstacionamientos retrieves all AdministradoresEstacionamientos matches certain condition. Returns empty list if
// no records exist
func GetAllAdministradoresEstacionamientos(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(AdministradoresEstacionamientos))
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

	var l []AdministradoresEstacionamientos
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

// UpdateAdministradoresEstacionamientos updates AdministradoresEstacionamientos by Id and returns error if
// the record to be updated doesn't exist
func UpdateAdministradoresEstacionamientosById(m *AdministradoresEstacionamientos) (err error) {
	o := orm.NewOrm()
	v := AdministradoresEstacionamientos{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteAdministradoresEstacionamientos deletes AdministradoresEstacionamientos by Id and returns error if
// the record to be deleted doesn't exist
func DeleteAdministradoresEstacionamientos(id int) (err error) {
	o := orm.NewOrm()
	v := AdministradoresEstacionamientos{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&AdministradoresEstacionamientos{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
