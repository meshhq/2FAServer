package main

import (
	"2FAServer/configuration"
	"2FAServer/db"
	"2FAServer/handler"
	"html/template"

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

	t := &handler.HTMLTemplate{
		Templates: template.Must(template.ParseGlob("public/views/otp.html")),
	}

	server.Renderer = t
	viewHandler := &handler.ViewHandler{}
	server.GET(configuration.APIPath+"/code", viewHandler.GetQRCode)

	server.Logger.Fatal(server.Start(":1323"))
}
