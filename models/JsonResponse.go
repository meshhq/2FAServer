package models

type JSONResponse struct {
	TimeStamp int64  `json:"timestamp"`
	Message   string `json:"message"`
}
