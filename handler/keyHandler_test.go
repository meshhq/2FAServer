package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"

	"2FAServer/configuration"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

var (
	h                     = new(KeyHandler)
	newKeyJSON            = `{"user_id":"aGithubUserId","key":"12345678901234567890","provider":"aProvider"}`
	newKeyJSONResponse    = `{"timestamp":` + strconv.Itoa(int(time.Now().Unix())) + `,"message":"12345678901234567890aProvideraGithubUserId"}`
	updateKeyJSONResponse = `{"timestamp":` + strconv.Itoa(int(time.Now().Unix())) + `,"message":"` + testUserID + `"}`
	deleteKeyJSONResponse = `{"timestamp":` + strconv.Itoa(int(time.Now().Unix())) + `,"message":"` + testUserID + `"}`
	testUserID            = "1e3243566776998723t3reververv"
)

func TestCreateKey(t *testing.T) {
	e := echo.New()

	req := httptest.NewRequest(echo.POST, configuration.APIPath, strings.NewReader(newKeyJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, h.CreateKey(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, newKeyJSONResponse, rec.Body.String())
	}
}

func TestGetKeys(t *testing.T) {
	e := echo.New()

	req := httptest.NewRequest(echo.GET, configuration.APIPath+"?user_id="+testUserID, nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, h.GetKeys(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		var raw []string
		json.Unmarshal(rec.Body.Bytes(), &raw)
		assert.True(t, len(raw) == 5)
	}
}

func TestUpdateKey(t *testing.T) {
	e := echo.New()

	req := httptest.NewRequest(echo.PUT, configuration.APIPath, nil)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.SetPath(configuration.APIPath + "/:key_id")
	c.SetParamNames("key_id")
	c.SetParamValues(testUserID)

	// Assertions
	if assert.NoError(t, h.UpdateKey(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		assert.Equal(t, updateKeyJSONResponse, rec.Body.String())
	}
}

func TestDeleteKey(t *testing.T) {
	e := echo.New()

	req := httptest.NewRequest(echo.DELETE, configuration.APIPath, nil)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.SetPath(configuration.APIPath + "/:key_id")
	c.SetParamNames("key_id")
	c.SetParamValues(testUserID)

	// Assertions
	if assert.NoError(t, h.DeleteKey(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		assert.Equal(t, deleteKeyJSONResponse, rec.Body.String())
	}
}
