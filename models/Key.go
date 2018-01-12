package models

import "fmt"

const collectionName = "keys"

type Key struct {
	KeyID     int64  `json:"key_id" validate:"required" sql:",pk"`
	Key       string `json:"key" validate:"required"`
	UserID    string `json:"user_id" validate:"required"`
	Provider  string `json:"provider" validate:"required"`
	CreatedAt int64  `json:"created_at"`
}

func (k Key) String() string {
	return fmt.Sprintf("%s:%s:%s", k.Key, k.UserID, k.Provider)
}

func (k Key) CollectionName() string {
	return collectionName
}

// ObjectID returns the object identifier.
func (k Key) ObjectID() int64 {
	return k.KeyID
}

// SetCreatedAt sets the created_at property of the model.
func (k Key) SetCreatedAt(timestamp int64) {
	k.CreatedAt = timestamp
}
