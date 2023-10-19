package feedback

import "github.com/dgraph-io/badger/v4"

type Manager struct {
	db *badger.DB
}

func (m *Manager) Save(f *Feedback) error {
	return nil
}

func (m *Manager) Fetch(id string) (*Feedback, error) {
	return nil, nil
}

func (m *Manager) FetchAll() ([]*Feedback, error) {
	return nil, nil
}

func New(db *badger.DB) *Manager {
	return &Manager{
		db: db,
	}
}

type Feedback struct {
}
