package store

import (
	"2FAServer/db"
	"2FAServer/models"
	"os"
	"testing"

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

func TestUpdateKey(t *testing.T) {
	nk := *testKey
	res, err := keyStore.UpdateKey(nk.ID, testUpdateKey.Key)

	// Assertions
	assert.Equal(t, nil, err)
	assert.Equal(t, true, res)
}

func TestUpdateKeyWithMissingID(t *testing.T) {
	nk := *testKey
	nk.ID = 0

	res, err := keyStore.UpdateKey(nk.ID, testUpdateKey.Key)

	// Assertions
	assert.NotEqual(t, nil, err)
	assert.Equal(t, false, res)
}

func TestDeleteKey(t *testing.T) {
	res, err := keyStore.DeleteKey(*testKey)

	// Assertions
	assert.Equal(t, nil, err)
	assert.Equal(t, true, res)
}

func TestDeleteKeyWithMissingID(t *testing.T) {
	nk := *testKey
	nk.ID = 0

	res, err := keyStore.DeleteKey(nk)

	// Assertions
	assert.NotEqual(t, nil, err)
	assert.Equal(t, false, res)
}

func TestKeyByID(t *testing.T) {
	k, err := keyStore.KeyByID(testKey.ID)

	// Assertions
	assert.Equal(t, nil, err)
	assert.Equal(t, testKey.ID, k.ID)
}

func TestKeyByIDWithMissingID(t *testing.T) {
	nk := *testKey
	nk.ID = 0

	k, err := keyStore.KeyByID(nk.ID)

	// Assertions
	assert.NotEqual(t, nil, err)
	assert.Equal(t, 0, int(k.ID))
}

func TestKeyByIDWithEmptyID(t *testing.T) {
	k, err := keyStore.KeyByID(0)

	// Assertions
	assert.NotEqual(t, nil, err)
	assert.Equal(t, 0, int(k.ID))
}

func TestKeysByUserID(t *testing.T) {
	keys, err := keyStore.KeysByUserID(testKey.UserID)

	// Assertions
	assert.Equal(t, nil, err)
	assert.Equal(t, 5, len(keys))
}

func TestKeysByUserIDWithMissingUserID(t *testing.T) {
	keys, err := keyStore.KeysByUserID("")

	// Assertions
	assert.NotEqual(t, nil, err)
	assert.Equal(t, 0, len(keys))
}
