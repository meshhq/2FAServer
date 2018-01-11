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
func (dbc *MockDbContext) GetKeyByID(keyID int) models.Key {
	return models.Key{KeyID: keyID}
}

// GetKeysByUserID retrieves a list of Keys by UserID
func (dbc *MockDbContext) GetKeysByUserID(userID string) []models.Key {
	return []models.Key{
		models.Key{KeyID: 1},
		models.Key{KeyID: 2},
		models.Key{KeyID: 3},
		models.Key{KeyID: 4},
		models.Key{KeyID: 5},
	}
}

// InsertKey creates a new Key record in the database.
func (dbc *MockDbContext) InsertKey(m models.Key) models.Key {
	if m.KeyID > 5 {
		return models.Key{}
	}

	return m
}

// UpdateKey updates a Key records's key value.
func (dbc *MockDbContext) UpdateKey(keyID int, key string) bool {
	if keyID < 5 {
		return true
	}

	return false
}

// DeleteKey removes a Key record from the database.
func (dbc *MockDbContext) DeleteKey(m models.Key) bool {
	if m.KeyID < 5 {
		return true
	}

	return false
}
