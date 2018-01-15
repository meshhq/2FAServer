package models

type Model struct {
	ID        int64 `json:"key_id,string" validate:"required" sql:",pk"`
	CreatedAt int64 `json:"created_at"`

	collectionName string
}

func (m *Model) CollectionName() string {
	return m.collectionName
}

// ObjectID returns the object identifier.
func (m *Model) ObjectID() int64 {
	return m.ID
}

// SetCreatedAt sets the created_at property of the model.
func (m *Model) SetCreatedAt(timestamp int64) {
	m.CreatedAt = timestamp
}

// SetCreatedAt sets the created_at property of the model.
func (m *Model) SetID(ID int64) {
	m.ID = ID
}
