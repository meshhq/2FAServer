package store

import (
	"os"
	"testing"

	"github.com/meshhq/2FAServer/db"
	"github.com/meshhq/2FAServer/models"

	"github.com/icrowley/fake"
	"github.com/stretchr/testify/assert"
)

var (
	database      db.ContextInterface = new(db.MockDbContext)
	keyStore                          = NewKeyStore(&database)
	testKey                           = new(models.Key)
	testUpdateKey                     = new(models.Key)
)

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

// Test InsertKey
func TestInsertKey(t *testing.T) {
	nk := *testKey
	nk.ID = 0

	k, err := keyStore.InsertKey(nk)

	// Assertions
	assert.Equal(t, nil, err)
	assert.Equal(t, int(testKey.ID), int(k.ID))
	assert.Equal(t, testKey.UserID, k.UserID)
	assert.Equal(t, testKey.Key, k.Key)
	assert.Equal(t, testKey.Provider, k.Provider)
}

// Test InsertKey With Existing ID in Model argument.
func TestInsertKeyWithExistingID(t *testing.T) {
	nk := *testKey
	nk.ID = 1

	k, err := keyStore.InsertKey(nk)

	// Assertions
	assert.NotEqual(t, nil, err)
	assert.Equal(t, 0, int(k.ID))
	assert.Equal(t, "", k.UserID)
	assert.Equal(t, "", k.Key)
	assert.Equal(t, "", k.Provider)
}

// Test InsertKey With Missing Provider
func TestInsertKeyWithMissingProvider(t *testing.T) {
	nk := *testKey
	nk.Provider = ""

	k, err := keyStore.InsertKey(nk)

	// Assertions
	assert.NotEqual(t, nil, err)
	assert.Equal(t, 0, int(k.ID))
	assert.Equal(t, "", k.UserID)
	assert.Equal(t, "", k.Key)
	assert.Equal(t, "", k.Provider)
}

// Test InsertKey With Missing User ID
func TestInsertKeyWithMissingUserID(t *testing.T) {
	nk := *testKey
	nk.UserID = ""

	k, err := keyStore.InsertKey(nk)

	// Assertions
	assert.NotEqual(t, nil, err)
	assert.Equal(t, 0, int(k.ID))
	assert.Equal(t, "", k.UserID)
	assert.Equal(t, "", k.Key)
	assert.Equal(t, "", k.Provider)
}

// Test InsertKey With Missing Key
func TestInsertKeyWithMissingKey(t *testing.T) {
	nk := *testKey
	nk.Key = ""

	k, err := keyStore.InsertKey(nk)

	// Assertions
	assert.NotEqual(t, nil, err)
	assert.Equal(t, 0, int(k.ID))
	assert.Equal(t, "", k.UserID)
	assert.Equal(t, "", k.Key)
	assert.Equal(t, "", k.Provider)
}

// TestUpdateKey
func TestUpdateKey(t *testing.T) {
	nk := *testKey
	res, err := keyStore.UpdateKey(nk.ID, testUpdateKey.Key)

	// Assertions
	assert.Equal(t, nil, err)
	assert.Equal(t, true, res)
}

// Test UpdateKey With Missing ID
func TestUpdateKeyWithMissingID(t *testing.T) {
	nk := *testKey
	nk.ID = 0

	res, err := keyStore.UpdateKey(nk.ID, testUpdateKey.Key)

	// Assertions
	assert.NotEqual(t, nil, err)
	assert.Equal(t, false, res)
}

// Test DeleteKey
func TestDeleteKey(t *testing.T) {
	res, err := keyStore.DeleteKey(*testKey)

	// Assertions
	assert.Equal(t, nil, err)
	assert.Equal(t, true, res)
}

// TestDeleteKeyWithMissingID
func TestDeleteKeyWithMissingID(t *testing.T) {
	nk := *testKey
	nk.ID = 0

	res, err := keyStore.DeleteKey(nk)

	// Assertions
	assert.NotEqual(t, nil, err)
	assert.Equal(t, false, res)
}

// TestKeyByID
func TestKeyByID(t *testing.T) {
	k, err := keyStore.KeyByID(testKey.ID)

	// Assertions
	assert.Equal(t, nil, err)
	assert.Equal(t, testKey.ID, k.ID)
}

// TestKeyByIDWithMissingID
func TestKeyByIDWithMissingID(t *testing.T) {
	nk := *testKey
	nk.ID = 0

	k, err := keyStore.KeyByID(nk.ID)

	// Assertions
	assert.NotEqual(t, nil, err)
	assert.Equal(t, 0, int(k.ID))
}

// TestKeyByIDWithEmptyID
func TestKeyByIDWithEmptyID(t *testing.T) {
	k, err := keyStore.KeyByID(0)

	// Assertions
	assert.NotEqual(t, nil, err)
	assert.Equal(t, 0, int(k.ID))
}

// Test KeysByUserID
func TestKeysByUserID(t *testing.T) {
	keys, err := keyStore.KeysByUserID(testKey.UserID)

	// Assertions
	assert.Equal(t, nil, err)
	assert.Equal(t, 5, len(keys))
}

// Test KeysByUserID With Missing UserID
func TestKeysByUserIDWithMissingUserID(t *testing.T) {
	keys, err := keyStore.KeysByUserID("")

	// Assertions
	assert.NotEqual(t, nil, err)
	assert.Equal(t, 0, len(keys))
}

// TestKeyByUserIDAndProvider
func TestKeyByUserIDAndProvider(t *testing.T) {
	key, err := keyStore.KeyByUserIDProvider(testKey.UserID, testKey.Provider)

	// Assertions
	assert.Equal(t, nil, err)
	assert.NotEqual(t, nil, key)
}

// Test KeyByUserIDAndProvider With Provider Missing
func TestKeyByUserIDAndProviderWithProviderMissing(t *testing.T) {
	key, err := keyStore.KeyByUserIDProvider(testKey.UserID, "")

	// Assertions
	assert.NotEqual(t, nil, err)
	assert.Equal(t, 0, int(key.ID))
}

// Test KeyByUserIDAndProvider With UserID Missing
func TestKeyByUserIDAndProviderWithUserIDMissing(t *testing.T) {
	key, err := keyStore.KeyByUserIDProvider("", testKey.Provider)

	// Assertions
	assert.NotEqual(t, nil, err)
	assert.Equal(t, 0, int(key.ID))
}
