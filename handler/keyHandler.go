package handler

import (
	"net/http"
	"time"

	"2FAServer/models"

	"github.com/labstack/echo"
)

type KeyHandler struct {
}

// Create new key
func (h *KeyHandler) CreateKey(c echo.Context) (err error) {
	nk := new(models.Key)

	if err = c.Bind(nk); err != nil {
		response := models.JSONResponse{Message: "Invalid request payload\n", TimeStamp: time.Now().Unix()}
		return c.JSON(http.StatusBadRequest, response)
	}

	// Search for key in db.
	// if it doesnt exist, create.

	// Return new value.
	response := models.JSONResponse{Message: nk.Key + nk.Provider + nk.UserID, TimeStamp: time.Now().Unix()}
	return c.JSON(http.StatusCreated, response)
}

// Retrieve all keys in storage
func (h *KeyHandler) GetKeys(c echo.Context) error {
	var userID = c.QueryParam("user_id")

	if userID == "" {
		err := models.JSONResponse{Message: "'user_id' is missing.", TimeStamp: time.Now().Unix()}
		return c.JSON(http.StatusBadRequest, err)
	}

	// Fetch keys from Db
	keys := [5]string{
		"bdfajfsnkjav",
		"bdfajfsnkjav",
		"bdfajfsnkjav",
		"bdfajfsnkjav",
		"bdfajfsnkjav"}

	return c.JSON(http.StatusOK, keys)
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
