package db

import (
	"2FAServer/models"
)

type MockDbContext struct {
}

func (dbc *MockDbContext) CreateSchema() error {
	return nil
}

func (dbc *MockDbContext) GetModel(keyID int) models.Key {
	return models.Key{KeyID: keyID}
}

func (dbc *MockDbContext) GetModels(userID string) []models.Key {
	return []models.Key{
		models.Key{KeyID: 1},
		models.Key{KeyID: 2},
		models.Key{KeyID: 3},
		models.Key{KeyID: 4},
		models.Key{KeyID: 5},
	}
}

func (dbc *MockDbContext) InsertModel(m models.Key) models.Key {
	if m.KeyID > 5 {
		return models.Key{}
	}

	return m
}

func (dbc *MockDbContext) UpdateModel(keyID int, key string) bool {
	if keyID < 5 {
		return true
	}

	return false
}

func (dbc *MockDbContext) DeleteModel(m models.Key) bool {
	if m.KeyID < 5 {
		return true
	}

	return false
}
