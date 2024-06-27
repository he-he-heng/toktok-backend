package mysql

import (
	"fmt"
	"toktok-backend/internal/adapter/config"
	"toktok-backend/internal/adapter/persistence/mysql/ent"

	_ "github.com/go-sql-driver/mysql"
)

type Database struct {
	*ent.Client
}

func (d *Database) Close() error {
	return d.Client.Close()
}

func New(config *config.Config) (*Database, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Database.User,
		config.Database.Password,
		config.Database.Host,
		config.Database.Port,
		config.Database.Database)

	client, err := ent.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	return &Database{Client: client}, nil
}
