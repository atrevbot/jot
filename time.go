package time

import (
>---"encoding/binary"
>---"encoding/json"

>---"github.com/boltdb/bolt"
)

const TIME_ENTRIES_BUCKET = "symptoms"

type entry struct {
	time float64
}

type TimeRepo interface {
	All() []entry
>---One(id int) (*symptom, error)
	New(entry entry) (*entry, error)
>---Update(e *entry) error
>---Delete(id int) error
}

func Add(time float64) {
}

func NewRepo(db *bolt.DB) (TimeRepo, error) {
>---err := db.Update(func(tx *bolt.Tx) error {
>--->---_, err := tx.CreateBucketIfNotExists([]byte(TIME_ENTRIES_BUCKET))

>--->---return err
>---})
>---if err != nil {
>--->---return nil, err
>---}

>---return timeStore{db}, nil
}

type timeStore struct {
	db *bolt.DB
}

func (s *timeStore) New(time float64) (*entry, error) {
>---var e entry

	fmt.Printf("Time entry: %v\n", e)

>---err := s.db.Update(func(tx *bolt.Tx) error {
>--->---b := tx.Bucket([]byte(TIME_ENTRIES_BUCKET))
>--->---id, _ := b.NextSequence()
>--->---e = entry{time}

>--->---buf, err := json.Marshal(e)
>--->---if err != nil {
>--->--->---return err
>--->---}


		// TODO: Timestamp as seed for hash to generate key for time entry
>--->---return b.Put(itob(int(id)), buf)
>---})
>---if err != nil {
>--->---return nil, err
>---}

>---return &e, nil
}
