package config

import (
	"sync"

	"github.com/spf13/viper"
)

var (
	once   sync.Once
	config *Config
)

type Config struct {
	Database `json:"database"`
}

type Database struct {
	Mysql Mysql `json:"mysql"`
}

type Mysql struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Database string `json:"database"`
	Password string `json:"password"`
	MaxOpen  int    `json:"maxOpen"`
	MaxIdle  int    `json:"maxIdle"`
}

func LoadConfig(filename string) *Config {
	once.Do(func() {
		viper.SetConfigType("yaml")
		viper.SetConfigFile(filename)

		err := viper.ReadInConfig()
		if err != nil {
			panic(err)
		}

		err = viper.Unmarshal(&config)
		if err != nil {
			panic(err)
		}

	})

	return config
}
