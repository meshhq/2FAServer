package store

import (
	"2FAServer/db"
	"2FAServer/models"
)

// NewKeyStore creates a new KeyStore with the supplied values
func NewKeyStore(database *db.ContextInterface) *KeyStore {
	key := new(models.Key)

	keyStore := new(KeyStore)
	keyStore.Database = *database
	keyStore.CollectionName = key.CollectionName()

	return keyStore
}

// KeyStore stores the connection ID and typed CRUD methods.
type KeyStore struct {
	Database       db.ContextInterface
	CollectionName string
}

// KeyByID retrieves a Key by its ID.
func (s *KeyStore) KeyByID(keyID int64) models.Key {
	aKey := new(models.Key)
	aKey.ID = keyID

	s.Database.GetModel(aKey)

	return *aKey
}

// KeysByUserID retrieves a list of Keys by UserID.
func (s *KeyStore) KeysByUserID(userID string) []models.Key {
	var keys []interface{}

	result := s.Database.GetWithWhere(new(models.Key), keys, "user_id = ?", userID)

	var keySlice []models.Key
	for _, record := range result {
		keySlice = append(keySlice, record.(models.Key))
	}

	return keySlice
}

// InsertKey creates a new Key record in the database.
func (s *KeyStore) InsertKey(key *models.Key) models.Key {
	newKeyID := s.Database.InsertModel(key)
	if newKeyID.ObjectID() == 0 {
		return models.Key{}
	}

	return *key
}

// UpdateKey updates a Key records's key value.
func (s *KeyStore) UpdateKey(keyID int64, key string) bool {
	aKey := new(models.Key)
	aKey.ID = keyID

	existingKey := s.Database.GetModel(aKey).(*models.Key)
	if existingKey.ObjectID() == 0 {
		return false
	}

	existingKey.Key = key

	return s.Database.UpdateModel(existingKey)
}

// DeleteKey removes a Key record from the database.
func (s *KeyStore) DeleteKey(key models.Key) bool {
	return s.Database.DeleteModel(&key)
}
