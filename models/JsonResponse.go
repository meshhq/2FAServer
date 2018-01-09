package models

import "time"

type JSONResponse struct {
	TimeStamp int64       `json:"timestamp"`
	Message   string      `json:"message,omitempty"`
	Data      interface{} `json:"data"`
}

func NewJSONResponse(data interface{}, message string) *JSONResponse {
	r := new(JSONResponse)
	r.TimeStamp = time.Now().Unix()
	r.Message = message
	r.Data = data

	return r
}
