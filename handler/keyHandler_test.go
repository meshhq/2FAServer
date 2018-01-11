package handler_test

import (
	"2FAServer/models"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"2FAServer/configuration"
	"2FAServer/db"
	"2FAServer/handler"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

var (
	h       = handler.KeyHandler{DbContext: new(db.MockDbContext)}
	testKey = models.Key{
		KeyID:    1,
		UserID:   "leonardojperez",
		Key:      "e890b5b83a133b70cea4b069f401baf4",
		Provider: "MeshStudioAuthProvider",
	}
	testUpdateKey = models.Key{
		Key: "f90721c90de9bd9ef516bea0b184fd30",
	}
)

func stringifyKey(k models.Key) string {
	nk, err := json.Marshal(k)
	if err != nil {
		panic("Error converting Key to JSON.")
	}

	res := string(nk)
	return res
}

func TestCreateKey(t *testing.T) {
	e := echo.New()

	payload := stringifyKey(testKey)
	req := httptest.NewRequest(echo.POST, configuration.APIPath, strings.NewReader(payload))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, h.CreateKey(c)) {
		expected, err := json.Marshal(models.NewJSONResponse(testKey, "Success"))
		if err != nil {
		}

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, string(expected), rec.Body.String())
	}
}

func TestGetKeys(t *testing.T) {
	e := echo.New()

	req := httptest.NewRequest(echo.GET, configuration.APIPath+"?user_id="+testKey.UserID, nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, h.GetKeys(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		res := models.JSONResponse{}
		json.Unmarshal(rec.Body.Bytes(), &res)

		assert.True(t, len(res.Data.([]interface{})) == 5)
	}
}

func TestUpdateKey(t *testing.T) {
	e := echo.New()

	payload := stringifyKey(testUpdateKey)
	req := httptest.NewRequest(echo.PUT, configuration.APIPath, strings.NewReader(payload))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.SetPath(configuration.APIPath + "/:key_id")
	c.SetParamNames("key_id")
	c.SetParamValues(strconv.Itoa(testKey.KeyID))

	// Assertions
	if assert.NoError(t, h.UpdateKey(c)) {
		expected, err := json.Marshal(models.NewJSONResponse(nil, "Success."))
		if err != nil {
		}

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, string(expected), rec.Body.String())
	}
}

func TestDeleteKey(t *testing.T) {
	e := echo.New()

	req := httptest.NewRequest(echo.DELETE, configuration.APIPath, nil)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.SetPath(configuration.APIPath + "/:key_id")
	c.SetParamNames("key_id")
	c.SetParamValues(strconv.Itoa(testKey.KeyID))

	// Assertions
	if assert.NoError(t, h.DeleteKey(c)) {
		expected, err := json.Marshal(models.NewJSONResponse(nil, "Success."))
		if err != nil {
		}

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, string(expected), rec.Body.String())
	}
}
