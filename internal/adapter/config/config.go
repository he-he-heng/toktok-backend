package config

import (
	"os"

	"github.com/pelletier/go-toml/v2"
)

type Config struct {
	Database struct {
		User     string
		Password string
		Host     string
		Port     string
		Database string
	}

	Token struct {
		Key             []byte
		AccessDuration  string
		RefreshDuration string
	}
}

func New(path string) (*Config, error) {
	config := new(Config)

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	err = toml.NewDecoder(file).Decode(config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
