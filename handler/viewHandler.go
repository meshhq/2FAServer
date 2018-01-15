package handler

import (
	"encoding/base64"
	"net/http"

	"github.com/labstack/echo"
	qrcode "github.com/skip2/go-qrcode"
)

// ViewHandler
type ViewHandler struct {
}

// GetQRCode
func (vh *ViewHandler) GetQRCode(c echo.Context) error {
	png, err := qrcode.Encode("https://reddit.com", qrcode.Medium, 200)
	if err != nil {
		return err
	}

	imgBase64Str := base64.StdEncoding.EncodeToString(png)

	c.Render(http.StatusOK, "qrcode", imgBase64Str)
	return nil
}
