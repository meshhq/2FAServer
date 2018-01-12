package store

import (
	"2FAServer/db"
	"2FAServer/models"
)

// NewUserStore creates a new UserStore with the supplied values
func NewKeyStore(database *db.ContextInterface) *KeyStore {
	key := new(models.Key)

	keyStore := new(KeyStore)
	keyStore.Database = *database
	keyStore.CollectionName = key.CollectionName()

	return keyStore
}

// Store models a collection for an individual model.
type KeyStore struct {
	Database       db.ContextInterface
	CollectionName string
}

// GetKeysByUserID retrieves a list of Keys by UserID
func (s *KeyStore) KeyByID(keyID int64) models.Key {
	aKey := models.Key{
		KeyID: keyID,
	}

	s.Database.GetModel(&aKey)

	return aKey
}

func (s *KeyStore) KeysByUserID(userID string) []models.Key {
	var keys []models.Key
	result := s.Database.GetWithWhere(&models.Key{}, "user_id = ?", userID)

	for _, record := range result {
		keys = append(keys, record.(models.Key))
	}

	return keys
}

// InsertKey creates a new Key record in the database.
func (s *KeyStore) InsertKey(key models.Key) models.Key {
	newKeyID := s.Database.InsertModel(&key)
	if newKeyID == 0 {
		return models.Key{}
	}

	return key
}

// UpdateKey updates a Key records's key value.
func (s *KeyStore) UpdateKey(keyID int64, key string) bool {
	aKey := models.Key{
		KeyID: keyID,
		Key:   key,
	}

	result := s.Database.UpdateModel(&aKey)
	return result
}

// DeleteKey removes a Key record from the database.
func (s *KeyStore) DeleteKey(model models.Key) bool {
	result := s.Database.DeleteModel(&model)
	return result
}
