package handler

import (
	"encoding/base64"
	"net/http"

	"2FAServer/configuration"
	"2FAServer/db"
	"2FAServer/models"
	"2FAServer/store"

	"github.com/labstack/echo"
	"github.com/pquerna/otp/totp"
	qrcode "github.com/skip2/go-qrcode"
)

// TOTPHandler Key Route handlers.
type TOTPHandler struct {
	store store.KeyStore
}

func NewTOTPHandler(database *db.ContextInterface) *TOTPHandler {
	totpHandler := new(TOTPHandler)

	store := store.NewKeyStore(database)
	totpHandler.store = *store

	return totpHandler
}

// Generate a new QR Code for 2FA enrollment.
func (h *TOTPHandler) Generate(c echo.Context) (err error) {
	requestKey := new(models.Key)
	if err = c.Bind(requestKey); err != nil {
		e := models.NewJSONResponse(nil, configuration.InvalidRequestPayload)
		return c.JSON(http.StatusBadRequest, e)
	}

	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      requestKey.Provider,
		AccountName: requestKey.UserID,
	})
	if err != nil {
		return err
	}

	// Store secret somewhere?
	savedKey := h.store.InsertKey(requestKey)
	if savedKey.ID == 0 {
		e := models.NewJSONResponse(nil, configuration.CreateKeyError)
		return c.JSON(http.StatusBadRequest, e)
	}

	png, err := qrcode.Encode(key.String(), qrcode.Medium, 200)
	if err != nil {
		return err
	}

	imgBase64Str := base64.StdEncoding.EncodeToString(png)
	c.Render(http.StatusOK, "qrcode", imgBase64Str)

	return nil
}

// Validate a supplied TOTP token.
func (h *TOTPHandler) Validate(c echo.Context) (err error) {
	token := new(models.Token)
	if err = c.Bind(token); err != nil {
		e := models.NewJSONResponse(nil, configuration.InvalidRequestPayload)
		return c.JSON(http.StatusBadRequest, e)
	}

	// Search for secret in storage by User and Provider

	valid := totp.Validate(token.Token, "")

	if valid {
		e := models.NewJSONResponse(nil, configuration.OTPValidationFailed)
		return c.JSON(http.StatusBadRequest, e)
	}

	return c.JSON(http.StatusBadRequest, models.NewJSONResponse(nil, configuration.Success))
}
