package db

import "2FAServer/models"

// DbContextInterface for DB access.
type DbContextInterface interface {
	CreateSchema() error
	GetModel(keyID int) models.Key
	GetModels(userID string) []models.Key
	InsertModel(m models.Key) models.Key
	UpdateModel(keyID int, key string) bool
	DeleteModel(m models.Key) bool
}
