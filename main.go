package main

import (
	"2FAServer/configuration"
	"2FAServer/handler"

	"github.com/labstack/echo"
)

func main() {
	server := echo.New()
	handler := new(handler.KeyHandler)

	server.POST(configuration.APIPath, handler.CreateKey)
	server.GET(configuration.APIPath, handler.GetKeys)
	server.PUT(configuration.APIPath+"/:key_id", handler.UpdateKey)
	server.DELETE(configuration.APIPath+"/:key_id", handler.DeleteKey)

	server.Logger.Fatal(server.Start(":1323"))
}
