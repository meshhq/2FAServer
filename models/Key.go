package models

import "fmt"

type Key struct {
	KeyID    string `json:"key_id" validate:"required" sql:",pk"`
	Key      string `json:"key" validate:"required"`
	UserID   string `json:"user_id" validate:"required"`
	Provider string `json:"provider" validate:"required"`
}

func (k Key) String() string {
	return fmt.Sprintf("%s:%s:%s", k.Key, k.UserID, k.Provider)
}
