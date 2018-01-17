package handler

import (
	"net/http"
	"strconv"

	"2FAServer/configuration"
	"2FAServer/db"
	"2FAServer/models"
	"2FAServer/store"

	"github.com/labstack/echo"
)

// KeyHandler Key Route handlers.
type KeyHandler struct {
	store store.KeyStore
}

func NewKeyHandler(database *db.ContextInterface) *KeyHandler {
	keyHandler := new(KeyHandler)

	store := store.NewKeyStore(database)
	keyHandler.store = *store

	return keyHandler
}

// CreateKey Creates new Key record.
func (h *KeyHandler) CreateKey(c echo.Context) (err error) {
	rk := new(models.Key)
	if err = c.Bind(rk); err != nil {
		return GetErrorResponse(c, configuration.InvalidRequestPayload)
	}

	nk, err := h.store.InsertKey(*rk)
	if err != nil || nk.ID == 0 {
		return GetErrorResponse(c, configuration.CreateKeyError)
	}

	return c.JSON(http.StatusOK, models.NewJSONResponse(nk, configuration.Success))
}

// GetKeys Retrieve all keys in storage by user_id
func (h *KeyHandler) GetKeys(c echo.Context) (err error) {
	var userID = c.QueryParam("user_id")
	if userID == "" {
		return GetErrorResponse(c, configuration.UserIDMissing)
	}

	keys, err := h.store.KeysByUserID(userID)
	if err != nil {
		return GetErrorResponse(c, configuration.KeysFetchError)
	}

	return c.JSON(http.StatusOK, models.NewJSONResponse(keys, configuration.Success))
}

// UpdateKey Updates existing key by key_id
func (h *KeyHandler) UpdateKey(c echo.Context) (err error) {
	keyID, err := strconv.ParseUint(c.Param("key_id"), 0, 0)
	if err != nil {
		return GetErrorResponse(c, configuration.KeyIDMissing)

	}

	payload := new(models.Key)
	if err = c.Bind(payload); err != nil {
		return GetErrorResponse(c, configuration.InvalidRequestPayload)
	}

	// Modify key property
	updated, err := h.store.UpdateKey(uint(keyID), payload.Key)
	if !updated || err != nil {
		return GetErrorResponse(c, configuration.UpdateKeyError)
	}

	return c.JSON(http.StatusOK, models.NewJSONResponse(nil, configuration.Success))
}

// DeleteKey Deletes existing key by Key_id
func (h *KeyHandler) DeleteKey(c echo.Context) error {
	keyID, err := strconv.ParseInt(c.Param("key_id"), 0, 0)
	if err != nil {
		return GetErrorResponse(c, configuration.InvalidRequestPayload)
	}

	if keyID == 0 {
		return GetErrorResponse(c, configuration.InvalidRequestPayload)
	}

	aKey := new(models.Key)
	aKey.ID = uint(keyID)

	removed, err := h.store.DeleteKey(*aKey)
	if !removed || err != nil {
		return GetErrorResponse(c, configuration.DeleteError)
	}

	return c.JSON(http.StatusOK, models.NewJSONResponse(nil, configuration.Success))
}
