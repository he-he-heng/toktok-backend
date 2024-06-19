package bolt

import (
	"io/fs"
	"time"

	"github.com/boltdb/bolt"
	"github.com/he-he-heng/toktok-backend/config"
)

const (
	StatusBucket string = "statusBucket"
)

type Database struct {
	*bolt.DB
}

func NewDatabase(c *config.Database) (*bolt.DB, error) {
	db, err := bolt.Open(c.Bolt.Path, fs.FileMode(c.Bolt.Mode), &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return nil, err
	}
	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(StatusBucket))
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return db, err
}

func (d *Database) Close() error {
	return d.DB.Close()
}
