package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

const collectionName = "keys"

type Key struct {
	gorm.Model
	Key      string `json:"key" validate:"required"`
	UserID   string `json:"user_id" validate:"required"`
	Provider string `json:"provider" validate:"required"`
}

// SetCreatedAt sets the created_at property of the model.
func (k *Key) SetCreatedAt(timestamp time.Time) {
	k.Model.CreatedAt = timestamp
}

// SetCreatedAt sets the created_at property of the model.
func (k *Key) SetID(ID uint) {
	k.Model.ID = ID
}
