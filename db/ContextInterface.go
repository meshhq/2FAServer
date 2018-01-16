package db

// ContextInterface for DB access.
type ContextInterface interface {
	GetModel(model interface{}) bool
	GetWithWhere(model interface{}, refArray interface{}, whereClause string, params ...interface{})
	InsertModel(model interface{}) bool
	UpdateModel(model interface{}) bool
	DeleteModel(model interface{}) bool
}
