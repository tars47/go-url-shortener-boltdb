package store

import (
	"fmt"

	"github.com/boltdb/bolt"
)

const DEFAULT_BUCKET_NAME = "url"

type Store struct {
	db *bolt.DB
}

func Connect() (Store, error) {
	conn, err := bolt.Open("mybolt.db", 0600, nil)

	if err != nil {
		return Store{}, err
	}

	db := Store{conn}

	if err := db.CreateBucket(); err != nil {
		return Store{}, err
	}

	return db, nil
}

func (s *Store) Close() {
	s.db.Close()
}

func (s *Store) CreateBucket() error {
	err := s.db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(DEFAULT_BUCKET_NAME))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) Upsert(path, url string) error {
	err := s.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(DEFAULT_BUCKET_NAME))
		err := b.Put([]byte(path), []byte(url))
		return err
	})
	return err
}

func (s *Store) Get(path string) (string, error) {
	var url string
	err := s.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(DEFAULT_BUCKET_NAME))
		urlbytes := b.Get([]byte(path))
		url = string(urlbytes)

		return nil
	})
	return url, err
}

func (s *Store) Delete(path string) error {

	err := s.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(DEFAULT_BUCKET_NAME))
		if err := b.Delete([]byte(path)); err != nil {
			return err
		}

		return nil
	})
	return err
}
