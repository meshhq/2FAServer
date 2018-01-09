package models

import "fmt"

type Key struct {
	Key      string `json:"key" validate:"required" sql:",pk"`
	UserID   string `json:"user_id" validate:"required"`
	Provider string `json:"provider" validate:"required"`
}

func (k Key) String() string {
	return fmt.Sprintf("%s:%s:%s", k.Key, k.UserID, k.Provider)
}
