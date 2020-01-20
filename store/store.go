package store

import (
	"encoding/json"
	"os"
	"time"

	bolt "go.etcd.io/bbolt"
)

const TIME_ENTRIES_BUCKET = "entries"

type entry struct {
	Created  time.Time
	Duration time.Duration
	Message  string
	Project  string
}

func NewEntry(d time.Duration, m string) (*entry, error) {
	p, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	return &entry{time.Now(), d, m, p}, nil
}

type Repo interface {
	All() ([]entry, error)
	One(id int) (*entry, error)
	Save(e *entry) error
	Update(e *entry) error
	Delete(id int) error
}

func NewRepo(db *bolt.DB) (Repo, error) {
	err := db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(TIME_ENTRIES_BUCKET))

		return err
	})
	if err != nil {
		return nil, err
	}

	return &store{db}, nil
}

type store struct {
	db *bolt.DB
}

func (s *store) All() ([]entry, error)      { var es []entry; return es, nil }
func (s *store) One(id int) (*entry, error) { return &entry{}, nil }
func (s *store) Save(e *entry) error {
	return s.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(TIME_ENTRIES_BUCKET))

		key, err := time.Now().MarshalText()
		if err != nil {
			return err
		}

		val, err := json.Marshal(e)
		if err != nil {
			return err
		}

		return b.Put(key, val)
	})
}
func (s *store) Update(e *entry) error { return nil }
func (s *store) Delete(id int) error   { return nil }
