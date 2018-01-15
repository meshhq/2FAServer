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
		e := models.NewJSONResponse(nil, configuration.InvalidRequestPayload)
		return c.JSON(http.StatusBadRequest, e)
	}

	nk := h.store.InsertKey(*rk)
	if nk.ID == 0 {
		e := models.NewJSONResponse(nil, configuration.Success)
		return c.JSON(http.StatusBadRequest, e)
	}

	return c.JSON(http.StatusOK, models.NewJSONResponse(nk, configuration.Success))
}

// GetKeys Retrieve all keys in storage by user_id
func (h *KeyHandler) GetKeys(c echo.Context) (err error) {
	var userID = c.QueryParam("user_id")
	if userID == "" {
		err := models.NewJSONResponse(nil, configuration.UserIDMissing)
		return c.JSON(http.StatusBadRequest, err)
	}

	keys := h.store.KeysByUserID(userID)
	return c.JSON(http.StatusOK, models.NewJSONResponse(keys, configuration.Success))
}

// UpdateKey Updates existing key by key_id
func (h *KeyHandler) UpdateKey(c echo.Context) (err error) {
	keyID, err := strconv.ParseInt(c.Param("key_id"), 0, 0)
	if err != nil {
		e := models.NewJSONResponse(nil, configuration.KeyIDMissing)
		return c.JSON(http.StatusBadRequest, e)
	}

	payload := new(models.Key)
	if err = c.Bind(payload); err != nil {
		e := models.NewJSONResponse(nil, configuration.InvalidRequestPayload)
		return c.JSON(http.StatusBadRequest, e)
	}

	// Modify key property
	updated := h.store.UpdateKey(keyID, payload.Key)
	if !updated {
		return c.JSON(http.StatusBadRequest, models.NewJSONResponse(nil, configuration.UpdateKeyError))
	}

	return c.JSON(http.StatusOK, models.NewJSONResponse(nil, configuration.Success))
}

// DeleteKey Deletes existing key by Key_id
func (h *KeyHandler) DeleteKey(c echo.Context) error {
	keyID, err := strconv.ParseInt(c.Param("key_id"), 0, 0)
	if err != nil {
		e := models.NewJSONResponse(nil, configuration.InvalidRequestPayload)
		return c.JSON(http.StatusBadRequest, e)
	}

	aKey := new(models.Key)
	aKey.ID = keyID

	removed := h.store.DeleteKey(*aKey)
	if !removed {
		return c.JSON(http.StatusBadRequest, models.NewJSONResponse(nil, configuration.DeleteError))
	}

	return c.JSON(http.StatusOK, models.NewJSONResponse(nil, configuration.Success))
}
