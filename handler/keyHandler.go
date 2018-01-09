package handler

import (
	"net/http"

	"2FAServer/db"
	"2FAServer/models"

	"github.com/labstack/echo"
)

var dbc = db.NewDbContext()

type KeyHandler struct {
}

// Create new Key record.
func (h *KeyHandler) CreateKey(c echo.Context) (err error) {
	nk := new(models.Key)
	if err = c.Bind(nk); err != nil {
		e := models.NewJSONResponse(nil, "Invalid request payload.")
		return c.JSON(http.StatusBadRequest, e)
	}

	// Retrieve existing if any and return it.
	existingKey := dbc.GetModel(nk.KeyID)
	if existingKey.KeyID != "" {
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
func (h *KeyHandler) GetKeys(c echo.Context) (err error) {
	var userID = c.QueryParam("user_id")
	if userID == "" {
		err := models.NewJSONResponse(nil, "Property 'user_id' is missing.")
		return c.JSON(http.StatusBadRequest, err)
	}

	keys := dbc.GetModels(userID)
	return c.JSON(http.StatusOK, models.NewJSONResponse(keys, ""))
}

// Update existing key by key_id
func (h *KeyHandler) UpdateKey(c echo.Context) (err error) {
	keyID := c.Param("key_id")
	key := new(models.Key)

	if err = c.Bind(key); err != nil {
		e := models.NewJSONResponse(nil, "Invalid request payload.")
		return c.JSON(http.StatusBadRequest, e)
	}

	// Search for key in db.
	existingKey := dbc.GetModel(keyID)
	if existingKey.Key == "" {
		return c.JSON(http.StatusBadRequest, models.NewJSONResponse(nil, "Element does not exist."))
	}

	// Modify key property
	updated := dbc.UpdateModel(keyID, key.Key)
	if !updated {
		return c.JSON(http.StatusBadRequest, models.NewJSONResponse(nil, "Could not update key."))
	}

	return c.JSON(http.StatusOK, models.NewJSONResponse(nil, keyID))
}

// Delete existing key by Key_id
func (h *KeyHandler) DeleteKey(c echo.Context) error {
	keyID := c.Param("key_id")

	removed := dbc.DeleteModel(models.Key{KeyID: keyID})
	if !removed {
		return c.JSON(http.StatusBadRequest, models.NewJSONResponse(nil, "Could not remove key."))
	}

	return c.JSON(http.StatusOK, models.NewJSONResponse(nil, keyID))
}
