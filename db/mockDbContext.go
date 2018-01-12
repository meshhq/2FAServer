package db

import (
	"2FAServer/models"
)

// MockDbContext defines a Mock object to use in unit tests.
type MockDbContext struct {
}

// CreateSchema defines a new of tables based on models.Key struct.
func (dbc *MockDbContext) CreateSchema() error {
	return nil
}

// GetKeyByID retrieves a model by KeyID
func (dbc *MockDbContext) GetModel(model models.Persistable) interface{} {
	return model.ObjectID()
}

// GetKeysByUserID retrieves a list of Keys by UserID
func (dbc *MockDbContext) GetWithWhere(model models.Persistable, whereClause string, params ...interface{}) []interface{} {
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
func (dbc *MockDbContext) InsertModel(model models.Persistable) int64 {
	if model.ObjectID() > 5 {
		return 0
	}

	return model.ObjectID()
}

// UpdateKey updates a Key records's key value.
func (dbc *MockDbContext) UpdateModel(model models.Persistable) bool {
	if model.ObjectID() < 5 {
		return true
	}

	return false
}

// DeleteKey removes a Key record from the database.
func (dbc *MockDbContext) DeleteModel(model models.Persistable) bool {
	if model.ObjectID() < 5 {
		return true
	}

	return false
}
