package handler

import (
	"net/http"

	"github.com/meshhq/2FAServer/models"

	"github.com/labstack/echo"
)

func GetErrorResponse(c echo.Context, message string) error {
	e := models.NewJSONResponse(nil, message)
	return c.JSON(http.StatusBadRequest, e)
}
