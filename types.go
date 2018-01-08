package main

type NewKey struct {
	UserID   string `json:"user_id" validate:"required"`
	Key      string `json:"key" validate:"required,key"`
	Provider string `json:"provider" validate:"required,provider"`
}

type JsonResponse struct {
	TimeStamp int64  `json:"timestamp"`
	Message   string `json:"message"`
}
