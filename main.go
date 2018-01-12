package main

import (
	"2FAServer/configuration"
	"2FAServer/db"
	"2FAServer/handler"

	"github.com/labstack/echo"
)

func main() {
	server := echo.New()
	database := db.NewDbContext()
	keyHandler := handler.NewKeyHandler(&database)

	server.POST(configuration.APIPath, keyHandler.CreateKey)
	server.GET(configuration.APIPath, keyHandler.GetKeys)
	server.PUT(configuration.APIPath+"/:key_id", keyHandler.UpdateKey)
	server.DELETE(configuration.APIPath+"/:key_id", keyHandler.DeleteKey)

	server.Logger.Fatal(server.Start(":1323"))
}
