package models

type Token struct {
	Token    string `json:"token" validate:"required"`
	UserID   string `json:"user_id" validate:"required"`
	Provider string `json:"provider" validate:"required"`
}
