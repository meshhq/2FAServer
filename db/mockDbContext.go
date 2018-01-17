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
func (dbc *MockDbContext) GetModel(model interface{}) bool {
	return true
}

// GetWithWhere retrieves a list of Models and filters by a where clause.
func (dbc *MockDbContext) GetWithWhere(refArray interface{}, whereClause string, params ...interface{}) {
	t := reflect.TypeOf(refArray)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	if t.Kind() == reflect.Slice {
		t = t.Elem()
	} else {
		panic("Input param is not a slice")
	}

	sl := reflect.ValueOf(refArray).Elem()
	if t.Kind() == reflect.Ptr {
		sl = sl.Elem()
	}

	sliceType := sl.Type().Elem()
	if sliceType.Kind() == reflect.Ptr {
		sliceType = sliceType.Elem()
	}

	for i := 0; i < 5; i++ {
		newitem := reflect.New(sliceType).Elem()
		sl.Set(reflect.Append(sl, newitem))
	}
}

// InsertModel creates a new Key record in the database.
func (dbc *MockDbContext) InsertModel(model interface{}) bool {
	nm := model.(*models.Key)
	nm.ID = 1

	return true
}

// UpdateModel updates a Key records's key value.
func (dbc *MockDbContext) UpdateModel(model interface{}) bool {
	return true
}

// DeleteModel removes a Key record from the database.
func (dbc *MockDbContext) DeleteModel(model interface{}) bool {
	return true
}
