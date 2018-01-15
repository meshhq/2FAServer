package store

import (
	"2FAServer/db"
	"2FAServer/models"
	"testing"

	"github.com/icrowley/fake"
	"github.com/stretchr/testify/assert"
)

var (
	database db.ContextInterface = new(db.MockDbContext)
	keyStore                     = NewKeyStore(&database)
	testKey                      = models.Key{
		ID:       1,
		UserID:   fake.UserName(),
		Key:      fake.Password(10, 20, true, true, true),
		Provider: fake.Word(),
	}
	testUpdateKey = models.Key{
		Key: fake.Password(10, 20, true, true, true),
	}
)

func TestInsertKey(t *testing.T) {
	nk := models.Key{
		UserID:   testKey.UserID,
		Key:      testKey.Key,
		Provider: testKey.Provider,
	}

	keyStore.InsertKey(&nk)

	// Assertions
	assert.Equal(t, testKey.ID, nk.ID)
	assert.Equal(t, testKey.UserID, nk.UserID)
	assert.Equal(t, testKey.Key, nk.Key)
	assert.Equal(t, testKey.Provider, nk.Provider)
}

func TestUpdateKey(t *testing.T) {
	nk := testKey
	res := keyStore.UpdateKey(nk.ID, testUpdateKey.Key)

	// Assertions
	assert.Equal(t, res, true)
}

func TestDeleteKey(t *testing.T) {
	res := keyStore.DeleteKey(testKey)

	// Assertions
	assert.Equal(t, res, true)
}

func TestKeyByID(t *testing.T) {
	k := keyStore.KeyByID(testKey.ID)

	// Assertions
	assert.Equal(t, testKey.ID, k.ID)
}

// // KeysByUserID retrieves a list of Keys by UserID.
// func (s *KeyStore) KeysByUserID(userID string) []models.Key {
// 	var keys []interface{}

// 	result := s.Database.GetWithWhere(new(models.Key), keys, "user_id = ?", userID)

// 	var keySlice []models.Key
// 	for _, record := range result {
// 		keySlice = append(keySlice, record.(models.Key))
// 	}

// 	return keySlice
// }
