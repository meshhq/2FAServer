package db

import (
	"2FAServer/models"
	"reflect"
)

// MockDbContext defines a Mock object to use in unit tests.
type MockDbContext struct {
}

// CreateSchema defines a new of tables based on models.Key struct.
func (dbc *MockDbContext) CreateSchema() error {
	return nil
}

// GetModel retrieves a model by PK.
func (dbc *MockDbContext) GetModel(model models.Persistable) models.Persistable {
	return model
}

// GetWithWhere retrieves a list of Models and filters by a where clause.
func (dbc *MockDbContext) GetWithWhere(model models.Persistable, refArray []interface{}, whereClause string, params ...interface{}) []interface{} {
	arr := []interface{}{
		&models.Model{ID: 1},
		&models.Model{ID: 2},
		&models.Model{ID: 3},
		&models.Model{ID: 4},
		&models.Model{ID: 5},
	}

	return arr
}

// InsertModel creates a new Key record in the database.
func (dbc *MockDbContext) InsertModel(model models.Persistable) models.Persistable {
	model.SetID(1)
	return model
}

// UpdateModel updates a Key records's key value.
func (dbc *MockDbContext) UpdateModel(model models.Persistable) bool {
	method := reflect.ValueOf(model).MethodByName("ObjectID")
	in := make([]reflect.Value, method.Type().NumIn())

	res := method.Call(in)
	id := res[0].Interface().(int64)

	if id < 5 {
		return true
	}

	return false
}

// DeleteModel removes a Key record from the database.
func (dbc *MockDbContext) DeleteModel(model models.Persistable) bool {
	method := reflect.ValueOf(model).MethodByName("ObjectID")
	in := make([]reflect.Value, method.Type().NumIn())

	res := method.Call(in)
	id := res[0].Interface().(int64)

	if id < 5 {
		return true
	}

	return false
}