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
	if err = c.Bind(&requestKey); err != nil {
		return GetErrorResponse(c, configuration.InvalidRequestPayload)
	}

	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      requestKey.Provider,
		AccountName: requestKey.UserID,
	})
	if err != nil {
		return GetErrorResponse(c, configuration.CreateOTPError)
	}

	requestKey.Key = key.Secret()

	// Store new key in db.
	savedKey, err := h.store.InsertKey(*requestKey)
	if err != nil || savedKey.ID == 0 {
		return GetErrorResponse(c, configuration.CreateKeyError)
	}

	png, err := qrcode.Encode(key.String(), qrcode.Medium, 200)
	if err != nil {
		return GetErrorResponse(c, configuration.CreateQRCodeError)
	}

	imgBase64Str := base64.StdEncoding.EncodeToString(png)
	c.Render(http.StatusOK, "qrcode", imgBase64Str)

	return nil
}

// Validate a supplied TOTP token.
func (h *TOTPHandler) Validate(c echo.Context) (err error) {
	token := new(models.Token)
	if err = c.Bind(token); err != nil {
		return GetErrorResponse(c, configuration.InvalidRequestPayload)
	}

	// TODO: Search for secret in storage by User and Provider
	key, err := h.store.KeysByUserID(token.UserID)
	if err != nil {
		return GetErrorResponse(c, configuration.OTPValidationFailed)
	}

	valid := totp.Validate(token.Token, key[0].Key)
	if valid {
		return GetErrorResponse(c, configuration.OTPValidationFailed)
	}

	return c.JSON(http.StatusOK, models.NewJSONResponse(nil, configuration.Success))
}
