package bolt

import (
	"io/fs"
	"time"

	"github.com/boltdb/bolt"
	"github.com/he-he-heng/toktok-backend/config"
)

const (
	PnumStatusBucket string = "pnumStatusBucket"
)

type Database struct {
	*bolt.DB
}

func NewDatabase(c *config.Config) (*Database, error) {
	db, err := bolt.Open(c.Database.Bolt.Path, fs.FileMode(c.Database.Bolt.Mode), &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return nil, err
	}
	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(PnumStatusBucket))
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return &Database{
		DB: db,
	}, err
}

func (d *Database) Close() error {
	return d.DB.Close()
}
