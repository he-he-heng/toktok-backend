package mysql

import (
	"fmt"
	"toktok-backend/internal/adapter/persistence/mysql/ent"
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

	ist, err := ent.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	client := Client{
		Client: ist,
	}

	return &client, nil
}

func (c *Client) Close() error {
	return c.Client.Close()
}
