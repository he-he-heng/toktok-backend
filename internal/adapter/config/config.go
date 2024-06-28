package config

import (
	"os"
	"toktok-backend/internal/core/domain"
	"toktok-backend/pkg/errors"

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
		Key             string
		AccessDuration  string
		RefreshDuration string
	}

	Handler struct {
		Port string
	}
}

func New(path string) (*Config, error) {
	config := new(Config)

	file, err := os.Open(path)
	if err != nil {
		return nil, errors.Wrap(domain.ErrInternal, err)
	}
	defer file.Close()

	err = toml.NewDecoder(file).Decode(config)
	if err != nil {
		return nil, errors.Wrap(domain.ErrInternal, err)
	}

	return config, nil
}
