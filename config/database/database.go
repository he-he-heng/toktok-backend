package database

import (
	"fmt"
	"time"

	"github.com/he-he-heng/toktok-backend/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	DefaultMaxOpenConns = 25
	DefaultMaxIdleConns = 25
)

type Database struct {
	db *gorm.DB
}

func NewDatabse(c *config.Config) *Database {
	db, err := gorm.Open(dialector(c.Database), nil)
	if err != nil {
		panic(err)
	}

	pool, err := db.DB()
	if err != nil {
		panic(err)
	}

	if c.Database.MaxOpen <= 0 {
		c.Database.MaxOpen = DefaultMaxOpenConns
	}
	if c.Database.MaxIdle <= 0 {
		c.Database.MaxIdle = DefaultMaxIdleConns
	}

	pool.SetMaxOpenConns(c.Database.MaxOpen)
	pool.SetMaxIdleConns(c.Database.MaxIdle)
	pool.SetConnMaxLifetime(5 * time.Second)

	return &Database{
		db: db,
	}
}

func dialector(c config.Database) gorm.Dialector {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		c.User,
		c.Password,
		c.Host,
		c.Port,
		c.Database,
	)

	return mysql.New(mysql.Config{
		DriverName:               "mysql",
		DSN:                      dsn,
		DefaultStringSize:        255,  // change it if needed
		DisableDatetimePrecision: true, // true, because datetime precision requires MySQL 5.6
		DontSupportRenameIndex:   true,
		DontSupportRenameColumn:  true,
	})
}
