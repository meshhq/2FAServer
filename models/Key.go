package models

type Key struct {
	UserID   string `json:"user_id" validate:"required"`
	Key      string `json:"key" validate:"required,key"`
	Provider string `json:"provider" validate:"required,provider"`
}
