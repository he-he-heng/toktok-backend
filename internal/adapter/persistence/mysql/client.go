package mysql

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"golang.org/x/exp/rand"

	"fmt"
	"toktok-backend/internal/adapter/persistence/mysql/ent"
	_ "toktok-backend/internal/adapter/persistence/mysql/ent/runtime"
	"toktok-backend/internal/infra/config"
)

type Client struct {
	*ent.Client
}

func NewClient(config *config.Config) (*Client, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Database.User,
		config.Database.Password,
		config.Database.Host,
		config.Database.Port,
		config.Database.Database)

	driver, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	// TODO: ping test
	if err := driver.DB().Ping(); err != nil {
		return nil, err
	}

	// TODO : max idle, conns setting

	maxPoolIdle := config.Database.MaxPoolIdle
	maxPoolOpen := config.Database.MaxPoolOpen
	maxPoolLifeTime := config.Database.MaxPoolLifeTime

	// // Get the underlying sql.DB object of the driver.
	db := driver.DB()
	db.SetMaxIdleConns(maxPoolIdle)
	db.SetMaxOpenConns(maxPoolOpen)
	db.SetConnMaxLifetime(time.Duration(rand.Int31n(int32(maxPoolLifeTime))) * time.Millisecond)
	ist := ent.NewClient(ent.Driver(driver))

	client := Client{
		Client: ist,
	}

	return &client, nil
}

func (c *Client) Close() error {
	return c.Client.Close()
}
