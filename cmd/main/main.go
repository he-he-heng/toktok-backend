package main

import (
	"fmt"

	"github.com/he-he-heng/toktok-backend/config"
)

func main() {
	cfg := config.LoadConfig(".env.yaml")
	fmt.Printf("%+v", cfg)

}
