package handler

import (
	"net/http"
	"time"

	"2FAServer/db"
	"2FAServer/models"

	"github.com/labstack/echo"
)

var dbc = db.NewDbContext()

type KeyHandler struct {
}

// Create new key
func (h *KeyHandler) CreateKey(c echo.Context) (err error) {
	nk := new(models.Key)
	if err = c.Bind(nk); err != nil {
		e := models.NewJSONResponse(nil, "Invalid request payload.")
		return c.JSON(http.StatusBadRequest, e)
	}

	// Retrieve existing if any.
	existingKey := dbc.GetModel(nk.Key)
	if existingKey.Key != "" {
		return c.JSON(http.StatusOK, models.NewJSONResponse(existingKey, ""))
	}

	saved := dbc.InsertModel(*nk)
	if !saved {
		e := models.NewJSONResponse(nil, "Could not create key.")
		return c.JSON(http.StatusBadRequest, e)
	}

	return c.JSON(http.StatusOK, models.NewJSONResponse(nk, ""))
}

// Retrieve all keys in storage by user_id
func (h *KeyHandler) GetKeys(c echo.Context) error {
	var userID = c.QueryParam("user_id")
	if userID == "" {
		err := models.NewJSONResponse(nil, "Property 'user_id' is missing.")
		return c.JSON(http.StatusBadRequest, err)
	}

	keys := dbc.GetModels(userID)
	return c.JSON(http.StatusOK, models.NewJSONResponse(keys, ""))
}

// Update existing key by key_id
func (h *KeyHandler) UpdateKey(c echo.Context) error {
	keyID := c.Param("key_id")

	// Search for key in db.
	// if it doesnt exist, do nothing

	// Modify key property

	response := models.JSONResponse{Message: keyID, TimeStamp: time.Now().Unix()}
	return c.JSON(http.StatusOK, response)
}

// Delete existing key by key_id
func (h *KeyHandler) DeleteKey(c echo.Context) error {
	keyID := c.Param("key_id")

	// Search for key in db.
	// if it doesnt exist, do nothing

	// Remove entry

	response := models.JSONResponse{Message: keyID, TimeStamp: time.Now().Unix()}
	return c.JSON(http.StatusOK, response)
}
