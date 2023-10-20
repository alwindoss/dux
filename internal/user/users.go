package user

import (
	"bytes"
	"encoding/gob"
	"log"
	"strings"

	"go.etcd.io/bbolt"
)

type Manager struct {
	db     *bbolt.DB
	bucket []byte
}

func (m *Manager) Save(u *User) error {
	err := m.db.Update(func(tx *bbolt.Tx) error {
		buck := tx.Bucket(m.bucket)
		key := []byte(u.Name)
		var valueBuff bytes.Buffer
		gob.NewEncoder(&valueBuff).Encode(u)
		value := valueBuff.Bytes()
		err := buck.Put(key, value)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		log.Printf("error saving user data with %s", u.ID)
	}
	log.Printf("saved user data with name %s successfully", u.Name)
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

	err := m.db.View(func(tx *bbolt.Tx) error {
		buck := tx.Bucket(m.bucket)
		c := buck.Cursor()
		for k, _ := c.First(); k != nil; k, _ = c.Next() {
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

func New(db *bbolt.DB) *Manager {

	m := &Manager{
		db:     db,
		bucket: []byte("users_bucket"),
	}
	db.Update(func(tx *bbolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(m.bucket)
		if err != nil {
			return err
		}
		return nil
	})
	return m
}

type User struct {
	ID   string
	Name string
}
