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

type Repo interface {
	All() ([]*entry, error)
	New(d time.Duration, m string) error
	Delete(id int) error
}

func New(db *bolt.DB) (Repo, error) {
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

func (s *store) All() ([]*entry, error) {
	var es []*entry

	err := s.db.View(func(tx *bolt.Tx) error {
		// Assume bucket exists and has keys
		b := tx.Bucket([]byte(TIME_ENTRIES_BUCKET))
		c := b.Cursor()

		// Query in descending order
		for k, v := c.Last(); k != nil; k, v = c.Prev() {
			e := &entry{}
			if err := json.Unmarshal(v, e); err != nil {
				return err
			}
			es = append(es, e)
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return es, nil
}
func (s *store) New(d time.Duration, m string) error {
	return s.db.Update(func(tx *bolt.Tx) error {
		key, err := time.Now().MarshalText()
		if err != nil {
			return err
		}

		p, err := os.Getwd()
		if err != nil {
			return err
		}

		val, err := json.Marshal(entry{time.Now(), d, m, p})
		if err != nil {
			return err
		}

		b := tx.Bucket([]byte(TIME_ENTRIES_BUCKET))

		return b.Put(key, val)
	})
}
func (s *store) Delete(id int) error { return nil }
