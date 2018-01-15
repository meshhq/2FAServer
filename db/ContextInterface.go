package db

import "2FAServer/models"

// ContextInterface for DB access.
type ContextInterface interface {
	CreateSchema() error
	GetModel(model models.Persistable) models.Persistable
	GetWithWhere(model models.Persistable, refArray []interface{}, whereClause string, params ...interface{}) []interface{}
	InsertModel(model models.Persistable) models.Persistable
	UpdateModel(model models.Persistable) bool
	DeleteModel(model models.Persistable) bool
}
