package main

import (
	"fmt"

	"github.com/he-he-heng/toktok-backend/config"
	database "github.com/he-he-heng/toktok-backend/config/database/mysql"
)

func main() {
	cfg := config.LoadConfig(".env.yaml")
	fmt.Printf("%+v", cfg)

	_, err := database.NewDatabse(cfg)
	if err != nil {
		panic(err)
	}
}
