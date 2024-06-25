package main

import (
	"fmt"
	"toktok-backend/internal/adapter/config"
)

func main() {
	config, err := config.New(".toml")
	if err != nil {
		panic(err)
	}
	fmt.Println(config)
}
