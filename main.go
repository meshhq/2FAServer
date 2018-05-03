package main

import (
	"html/template"

	"github.com/meshhq/2FAServer/configuration"
	"github.com/meshhq/2FAServer/db"
	"github.com/meshhq/2FAServer/handler"
	"github.com/meshhq/2FAServer/qrcode"

	"github.com/labstack/echo"
)

func main() {
	server := echo.New()

	database := db.NewDbContext()
	keyHandler := handler.NewKeyHandler(database)

	server.POST(configuration.KeysAPIPath, keyHandler.CreateKey)
	server.GET(configuration.KeysAPIPath, keyHandler.GetKeys)
	server.PUT(configuration.KeysAPIPath+"/:key_id", keyHandler.UpdateKey)
	server.DELETE(configuration.KeysAPIPath+"/:key_id", keyHandler.DeleteKey)

	data, err := qrcode.Asset("public/views/otp.html")
	if err != nil {
		panic(err)
	}

	optTemplate := template.New("OptTemplate")
	optTemplate.Parse(string(data))
	t := &handler.HTMLTemplate{
		Templates: template.Must(optTemplate, nil),
	}

	server.Renderer = t
	otpHandler := handler.NewTOTPHandler(&database)
	server.POST(configuration.OtpAPIPath, otpHandler.Generate)
	server.POST(configuration.OtpAPIPath+"/validate", otpHandler.Validate)

	server.Logger.Fatal(server.Start(":1323"))
}
