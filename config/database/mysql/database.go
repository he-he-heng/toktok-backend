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
	*gorm.DB
}

func NewDatabse(c *config.Config) (*Database, error) {
	db, err := gorm.Open(dialector(c.Database), nil)
	if err != nil {
		return nil, err
	}

	pool, err := db.DB()
	if err != nil {
		return nil, err
	}

	if c.Database.Mysql.MaxOpen <= 0 {
		c.Database.Mysql.MaxOpen = DefaultMaxOpenConns
	}
	if c.Database.Mysql.MaxIdle <= 0 {
		c.Database.Mysql.MaxIdle = DefaultMaxIdleConns
	}

	pool.SetMaxOpenConns(c.Database.Mysql.MaxOpen)
	pool.SetMaxIdleConns(c.Database.Mysql.MaxIdle)
	pool.SetConnMaxLifetime(5 * time.Second)

	return &Database{
		DB: db,
	}, nil
}

func dialector(c config.Database) gorm.Dialector {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		c.Mysql.User,
		c.Mysql.Password,
		c.Mysql.Host,
		c.Mysql.Port,
		c.Mysql.Database,
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
