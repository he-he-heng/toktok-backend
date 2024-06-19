package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Database Database `json:"database"`
}

type Database struct {
	Mysql Mysql `json:"mysql"`
	Bolt  Bolt  `json:"bolt"`
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

type Bolt struct {
	Path string `json:"path"`
	Mode int    `json:"mode"`
}

func NewConfig(filename string) (*Config, error) {
	viper.SetConfigType("yaml")
	viper.SetConfigFile(filename)

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	config := &Config{}
	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
