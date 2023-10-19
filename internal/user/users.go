package user

import (
	"bytes"
	"encoding/gob"
	"log"
	"strings"

	"github.com/dgraph-io/badger/v4"
)

type Manager struct {
	db *badger.DB
}

func (m *Manager) Save(u *User) error {

	err := m.db.Update(func(txn *badger.Txn) error {
		key := []byte(u.Name)
		var valueBuff bytes.Buffer
		gob.NewEncoder(&valueBuff).Encode(u)
		value := valueBuff.Bytes()
		err := txn.Set(key, value)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		log.Printf("error saving user data with %s", u.ID)
	}

	// log.Printf("saved user data with %s successfully", u.ID)
	return nil
}

func (m *Manager) Fetch(id string) (*User, error) {
	return nil, nil
}

func (m *Manager) FetchAll() ([]*User, error) {
	return nil, nil
}

func (m *Manager) Query(q string) ([]string, error) {
	userList := []string{}
	err := m.db.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		opts.PrefetchSize = 10
		it := txn.NewIterator(opts)
		defer it.Close()
		for it.Rewind(); it.Valid(); it.Next() {
			item := it.Item()
			k := item.Key()
			name := string(k)
			if strings.Contains(name, q) {
				userList = append(userList, name)
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return userList, nil
}

func New(db *badger.DB) *Manager {
	m := &Manager{
		db: db,
	}
	return m
}

type User struct {
	ID   string
	Name string
}
