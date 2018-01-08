package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

var (
	h                  = new(Handler)
	newKeyJSON         = `{"user_id":"aGithubUserId","key":"12345678901234567890","provider":"aProvider"}`
	newKeyJSONResponse = `{"timestamp":` + strconv.Itoa(int(time.Now().Unix())) + `,"message":"12345678901234567890aProvideraGithubUserId"}`
	testUserID         = "1e3243566776998723t3reververv"
)

func createRequest(verb string, payload string, path string) *http.Request {
	var req *http.Request

	switch verb {
	case echo.GET:
		req = httptest.NewRequest(verb, ApiPath+path, nil)
		break

	default:
		req = httptest.NewRequest(verb, ApiPath+path, strings.NewReader(payload))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	}

	return req
}

func TestCreateKey(t *testing.T) {
	// Setup
	e := echo.New()
	req := createRequest(echo.POST, newKeyJSON, "")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, h.createKey(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, newKeyJSONResponse, rec.Body.String())
	}
}

func TestGetKeys(t *testing.T) {
	// Setup
	e := echo.New()

	req := createRequest(echo.GET, "", "?user_id="+testUserID)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, h.getKeys(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		var raw []string
		json.Unmarshal(rec.Body.Bytes(), &raw)
		assert.True(t, len(raw) == 5)
	}
}
