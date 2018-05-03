package handler_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/icrowley/fake"

	"github.com/meshhq/2FAServer/configuration"
	"github.com/meshhq/2FAServer/db"
	"github.com/meshhq/2FAServer/handler"
	"github.com/meshhq/2FAServer/models"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

var (
	database      db.ContextInterface = new(db.MockDbContext)
	h                                 = handler.NewKeyHandler(database)
	testKey                           = new(models.Key)
	testUpdateKey                     = new(models.Key)
)

func stringifyKey(k models.Key) string {
	nk, err := json.Marshal(&k)
	if err != nil {
		panic("Error converting Key to JSON.")
	}

	res := string(nk)
	return res
}

func setup() {
	testKey.ID = 1
	testKey.UserID = fake.UserName()
	testKey.Key = fake.Password(10, 20, true, true, true)
	testKey.Provider = fake.Word()

	testUpdateKey.Key = fake.Password(10, 20, true, true, true)
}

func TestMain(m *testing.M) {
	setup()
	retCode := m.Run()
	os.Exit(retCode)
}

func TestCreateKey(t *testing.T) {
	e := echo.New()

	nk := &models.Key{
		Key:      testKey.Key,
		Provider: testKey.Provider,
		UserID:   testKey.UserID,
	}

	payload := stringifyKey(*nk)
	req := httptest.NewRequest(echo.POST, configuration.KeysAPIPath, strings.NewReader(payload))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, h.CreateKey(c)) {
		expected, err := json.Marshal(models.NewJSONResponse(testKey, configuration.Success))
		if err != nil {
		}

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, string(expected), rec.Body.String())
	}
}

func TestCreateKeyWithMissingArguments(t *testing.T) {
	e := echo.New()

	nk := &models.Key{
		Key:    testKey.Key,
		UserID: testKey.UserID,
	}

	payload := stringifyKey(*nk)
	req := httptest.NewRequest(echo.POST, configuration.KeysAPIPath, strings.NewReader(payload))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, h.CreateKey(c)) {
		expected, err := json.Marshal(models.NewJSONResponse(nil, configuration.CreateKeyError))
		if err != nil {
		}

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Equal(t, string(expected), rec.Body.String())
	}
}

func TestGetKeys(t *testing.T) {
	e := echo.New()

	req := httptest.NewRequest(echo.GET, configuration.KeysAPIPath+"?user_id="+testKey.UserID, nil)
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

	payload := stringifyKey(*testUpdateKey)
	req := httptest.NewRequest(echo.PUT, configuration.KeysAPIPath, strings.NewReader(payload))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.SetPath(configuration.KeysAPIPath + "/:key_id")
	c.SetParamNames("key_id")
	c.SetParamValues(strconv.Itoa(int(testKey.ID)))

	// Assertions
	if assert.NoError(t, h.UpdateKey(c)) {
		expected, err := json.Marshal(models.NewJSONResponse(nil, configuration.Success))
		if err != nil {
		}

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, string(expected), rec.Body.String())
	}
}

func TestDeleteKey(t *testing.T) {
	e := echo.New()

	req := httptest.NewRequest(echo.DELETE, configuration.KeysAPIPath, nil)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.SetPath(configuration.KeysAPIPath + "/:key_id")
	c.SetParamNames("key_id")
	c.SetParamValues(strconv.Itoa(int(testKey.ID)))

	// Assertions
	if assert.NoError(t, h.DeleteKey(c)) {
		expected, err := json.Marshal(models.NewJSONResponse(nil, "Success."))
		if err != nil {
		}

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, string(expected), rec.Body.String())
	}
}
