package mysql

import (
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

	ist, err := ent.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	// TODO: ping test
	// if err := drv.DB().Ping(); err != nil {
	// 	return nil, err
	// }

	// TODO : max idle, conns setting

	// maxPoolIdle, err := strconv.Atoi(os.Getenv("DB_POOL_IDLE_CONN"))
	// maxPoolOpen, err := strconv.Atoi(os.Getenv("DB_POOL_MAX_CONN"))
	// maxPollLifeTime, err := strconv.Atoi(os.Getenv("DB_POOL_LIFE_TIME"))
	// if err != nil {
	// 	return nil, err
	// }

	// // Get the underlying sql.DB object of the driver.
	// db := drv.DB()
	// db.SetMaxIdleConns(maxPoolIdle)
	// db.SetMaxOpenConns(maxPoolOpen)
	// db.SetConnMaxLifetime(time.Duration(rand.Int31n(int32(maxPollLifeTime))) * time.Millisecond)
	// clnt := ent.NewClient(ent.Driver(drv))

	client := Client{
		Client: ist,
	}

	return &client, nil
}

func (c *Client) Close() error {
	return c.Client.Close()
}
