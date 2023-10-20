package feedback

import (
	"go.etcd.io/bbolt"
)

type Manager struct {
	db *bbolt.DB
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

func New(db *bbolt.DB) *Manager {
	return &Manager{
		db: db,
	}
}

type Feedback struct {
}
