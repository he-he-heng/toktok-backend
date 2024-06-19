package config

import "sync"

var (
	once   sync.Once
	config *Config
)

type Config struct {
	Database `json:"database"`
}

type Database struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Database string `json:"database"`
	Password string `json:"password"`
	Type     string `json:"type"`
	MaxOpen  int    `json:"maxOpen"`
	MaxIdle  int    `json:"maxIdle"`
}
