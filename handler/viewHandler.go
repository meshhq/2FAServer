package handler

import (
	"github.com/labstack/echo"
)

// ViewHandler
type ViewHandler struct {
}

// GetQRCode
func (vh *ViewHandler) GetQRCode(c echo.Context) {
	// Return opt in barcode form.
}
