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
	return bolt.Open(c.Bolt.Path, fs.FileMode(c.Bolt.Mode), &bolt.Options{Timeout: 1 * time.Second})
}

func (d *Database) Close() error {
	return d.DB.Close()
}
