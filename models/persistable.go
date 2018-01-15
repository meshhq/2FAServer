package models

// Persistable declares an interface to which all types that will be used for
// database operations should conform.
type Persistable interface {

	// CollectionName returns a string representing the NoSQL collection
	// name for the type.
	CollectionName() string

	// ObjectID returns the object identifier for the type.
	ObjectID() int64

	// Sets the object identifier for the type.
	SetID(int64)

	// SetCreatedAt hydrates the created_at property of the model.
	SetCreatedAt(int64)
}
