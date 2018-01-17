package store

import (
	"2FAServer/configuration"
	"2FAServer/db"
	"2FAServer/models"
	"errors"
)

// NewKeyStore creates a new KeyStore with the supplied values
func NewKeyStore(database *db.ContextInterface) *KeyStore {
	keyStore := new(KeyStore)
	keyStore.Database = *database

	return keyStore
}

// KeyStore stores the connection ID and typed CRUD methods.
type KeyStore struct {
	Database       db.ContextInterface
	CollectionName string
}

// KeyByID retrieves a Key by its ID.
func (s *KeyStore) KeyByID(keyID uint) (models.Key, error) {
	if keyID == 0 {
		return models.Key{}, errors.New(configuration.KeyIDMustBeEmtpy)
	}

	aKey := &models.Key{}
	aKey.ID = keyID

	s.Database.GetModel(&aKey)

	return *aKey, nil
}

// KeysByUserID retrieves a list of Keys by UserID.
func (s *KeyStore) KeysByUserID(userID string) ([]models.Key, error) {
	var keys []models.Key
	if userID == "" {
		return keys, errors.New(configuration.UserIDMissing)
	}

	s.Database.GetWithWhere(&models.Key{}, &keys, "user_id = ?", userID)

	var keySlice []models.Key
	for _, record := range keys {
		keySlice = append(keySlice, record)
	}

	if len(keySlice) == 0 {
		keySlice = []models.Key{}
	}

	return keySlice, nil
}

// InsertKey creates a new Key record in the database.
func (s *KeyStore) InsertKey(key models.Key) (models.Key, error) {
	if key.ID != 0 {
		return models.Key{}, errors.New(configuration.KeyIDMustBeEmtpy)
	}

	if key.Key == "" {
		return models.Key{}, errors.New(configuration.KeySecretMissing)
	}

	if key.Provider == "" {
		return models.Key{}, errors.New(configuration.ProviderMissing)
	}

	s.Database.InsertModel(&key)
	if key.ID == 0 {
		return models.Key{}, errors.New(configuration.CreateKeyError)
	}

	return key, nil
}

// UpdateKey updates a Key records's key value.
func (s *KeyStore) UpdateKey(keyID uint, key string) (bool, error) {
	if keyID == 0 {
		return false, errors.New(configuration.KeyIDMissing)
	}

	if key == "" {
		return false, errors.New(configuration.KeySecretMissing)
	}

	aKey := new(models.Key)
	aKey.ID = keyID

	s.Database.GetModel(aKey)
	if aKey.ID == 0 {
		return false, errors.New(configuration.UpdateKeyError)
	}

	aKey.Key = key

	return s.Database.UpdateModel(aKey), nil
}

// DeleteKey removes a Key record from the database.
func (s *KeyStore) DeleteKey(key models.Key) (bool, error) {
	if key.ID == 0 {
		return false, errors.New(configuration.KeyIDMissing)
	}

	return s.Database.DeleteModel(&key), nil
}
