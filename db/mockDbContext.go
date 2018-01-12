package db

import (
	"2FAServer/models"
	"fmt"
	"reflect"
)

// MockDbContext defines a Mock object to use in unit tests.
type MockDbContext struct {
}

// CreateSchema defines a new of tables based on models.Key struct.
func (dbc *MockDbContext) CreateSchema() error {
	return nil
}

// GetKeyByID retrieves a model by KeyID
func (dbc *MockDbContext) GetModel(model interface{}) interface{} {
	return model
}

// GetKeysByUserID retrieves a list of Keys by UserID
func (dbc *MockDbContext) GetWithWhere(model interface{}, whereClause string, params ...interface{}) []interface{} {
	var arr []interface{}
	intList := [5]models.Persistable{
		models.Key{},
		models.Key{},
		models.Key{},
		models.Key{},
		models.Key{},
	}

	for _, val := range intList {
		arr = append(arr, val)
	}

	return arr
}

// InsertKey creates a new Key record in the database.
func (dbc *MockDbContext) InsertModel(model interface{}) interface{} {
	// pointer to struct - addressable
	ps := reflect.ValueOf(&model)
	// struct
	s := ps.Elem()
	if s.Kind() == reflect.Struct {
		// exported field
		f := s.FieldByName("KeyID")
		if f.IsValid() && f.CanSet() {
			// A Value can be changed only if it is
			// addressable and was not obtained by
			// the use of unexported struct fields.

			// change value of model.KeyID
			if f.Kind() == reflect.Int {
				x := int64(7)
				if !f.OverflowInt(x) {
					f.SetInt(x)
				}
			}

		}
	}

	// N at end
	fmt.Println(model)

	method := reflect.ValueOf(model).MethodByName("ObjectID")
	in := make([]reflect.Value, method.Type().NumIn())

	res := method.Call(in)
	id := res[0].Interface().(int64)

	if id < 5 {
		return true
	}

	return false
}

// UpdateKey updates a Key records's key value.
func (dbc *MockDbContext) UpdateModel(model interface{}) bool {
	method := reflect.ValueOf(model).MethodByName("ObjectID")
	in := make([]reflect.Value, method.Type().NumIn())

	res := method.Call(in)
	id := res[0].Interface().(int64)

	if id < 5 {
		return true
	}

	return false
}

// DeleteKey removes a Key record from the database.
func (dbc *MockDbContext) DeleteModel(model interface{}) bool {
	method := reflect.ValueOf(model).MethodByName("ObjectID")
	in := make([]reflect.Value, method.Type().NumIn())

	res := method.Call(in)
	id := res[0].Interface().(int64)

	if id < 5 {
		return true
	}

	return false
}
