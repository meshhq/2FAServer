package main

import (
	"github.com/labstack/echo"
)

func main() {
	server := echo.New()
	handler := new(Handler)

	server.POST("/keys", handler.createKey)
	server.GET("/keys", handler.getKeys)
	server.PUT("/keys/:key_id", handler.updateKey)
	server.DELETE("/keys/:key_id", handler.deleteKey)

	server.Logger.Fatal(server.Start(":1323"))
}
