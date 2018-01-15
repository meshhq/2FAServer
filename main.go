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

	server.POST(configuration.KeysAPIPath, keyHandler.CreateKey)
	server.GET(configuration.KeysAPIPath, keyHandler.GetKeys)
	server.PUT(configuration.KeysAPIPath+"/:key_id", keyHandler.UpdateKey)
	server.DELETE(configuration.KeysAPIPath+"/:key_id", keyHandler.DeleteKey)

	t := &handler.HTMLTemplate{
		Templates: template.Must(template.ParseGlob("public/views/otp.html")),
	}

	server.Renderer = t
	otpHandler := &handler.TOTPHandler{}
	server.POST(configuration.OtpAPIPath+"/code", otpHandler.Generate)
	server.POST(configuration.OtpAPIPath+"/code/validate", otpHandler.Validate)

	server.Logger.Fatal(server.Start(":1323"))
}
