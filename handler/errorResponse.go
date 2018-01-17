package handler

import (
	"2FAServer/models"
	"net/http"

	"github.com/labstack/echo"
)

func GetErrorResponse(c echo.Context, message string) error {
	e := models.NewJSONResponse(nil, message)
	return c.JSON(http.StatusBadRequest, e)
}
