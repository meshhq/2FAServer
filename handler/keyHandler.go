package handler

import (
	"net/http"
	"strconv"

	"2FAServer/db"
	"2FAServer/models"

	"github.com/labstack/echo"
)

// KeyHandler Key Route handlers.
type KeyHandler struct {
	DbContext db.DbContextInterface
}

// CreateKey Creates new Key record.
func (h *KeyHandler) CreateKey(c echo.Context) (err error) {
	rk := new(models.Key)
	if err = c.Bind(rk); err != nil {
		e := models.NewJSONResponse(nil, "Invalid request payload.")
		return c.JSON(http.StatusBadRequest, e)
	}

	nk := h.DbContext.InsertModel(*rk)
	if nk.KeyID == 0 {
		e := models.NewJSONResponse(nil, "Could not create key.")
		return c.JSON(http.StatusBadRequest, e)
	}

	return c.JSON(http.StatusOK, models.NewJSONResponse(nk, "Success"))
}

// GetKeys Retrieve all keys in storage by user_id
func (h *KeyHandler) GetKeys(c echo.Context) (err error) {
	var userID = c.QueryParam("user_id")
	if userID == "" {
		err := models.NewJSONResponse(nil, "Property 'user_id' is missing.")
		return c.JSON(http.StatusBadRequest, err)
	}

	keys := h.DbContext.GetModels(userID)
	return c.JSON(http.StatusOK, models.NewJSONResponse(keys, "Success"))
}

// UpdateKey Updates existing key by key_id
func (h *KeyHandler) UpdateKey(c echo.Context) (err error) {
	keyID, err := strconv.Atoi(c.Param("key_id"))
	if err != nil {
		e := models.NewJSONResponse(nil, "Invalid request payload. KeyID is missing.")
		return c.JSON(http.StatusBadRequest, e)
	}

	payload := new(models.Key)
	if err = c.Bind(payload); err != nil {
		e := models.NewJSONResponse(nil, "Invalid request payload.")
		return c.JSON(http.StatusBadRequest, e)
	}

	// Search for key in db.
	existingKey := h.DbContext.GetModel(keyID)
	if existingKey.KeyID == 0 {
		return c.JSON(http.StatusBadRequest, models.NewJSONResponse(nil, "Element does not exist."))
	}

	// Modify key property
	updated := h.DbContext.UpdateModel(keyID, payload.Key)
	if !updated {
		return c.JSON(http.StatusBadRequest, models.NewJSONResponse(nil, "Could not update key."))
	}

	return c.JSON(http.StatusOK, models.NewJSONResponse(nil, "Success."))
}

// DeleteKey Deletes existing key by Key_id
func (h *KeyHandler) DeleteKey(c echo.Context) error {
	keyID, err := strconv.Atoi(c.Param("key_id"))
	if err != nil {
		e := models.NewJSONResponse(nil, "Invalid request payload.")
		return c.JSON(http.StatusBadRequest, e)
	}

	removed := h.DbContext.DeleteModel(models.Key{KeyID: keyID})
	if !removed {
		return c.JSON(http.StatusBadRequest, models.NewJSONResponse(nil, "Could not remove key."))
	}

	return c.JSON(http.StatusOK, models.NewJSONResponse(nil, "Success."))
}
