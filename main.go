package main

import (
	"github.com/labstack/echo"
)

const ApiPath = "/keys"

func main() {
	server := echo.New()
	handler := new(Handler)

	server.POST(ApiPath, handler.createKey)
	server.GET(ApiPath, handler.getKeys)
	server.PUT(ApiPath+"/:key_id", handler.updateKey)
	server.DELETE(ApiPath+"/:key_id", handler.deleteKey)

	server.Logger.Fatal(server.Start(":1323"))
}
