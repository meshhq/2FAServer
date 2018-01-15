package models

const collectionName = "keys"

type Key struct {
	ID       int64  `json:"key_id,string" validate:"required" sql:",pk"`
	Key      string `json:"key" validate:"required"`
	UserID   string `json:"user_id" validate:"required"`
	Provider string `json:"provider" validate:"required"`
}

// CollectionName return the name of the Model's table.
func (k *Key) CollectionName() string {
	return collectionName
}

// ObjectID returns the object identifier.
func (k *Key) ObjectID() int64 {
	return k.ID
}

// SetCreatedAt sets the created_at property of the model.
func (k *Key) SetCreatedAt(timestamp int64) {
	//k.CreatedAt = timestamp
}

// SetCreatedAt sets the created_at property of the model.
func (k *Key) SetID(ID int64) {
	k.ID = ID
}
