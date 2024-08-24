package main

import (
	"context"
	"log"
	"toktok-backend/internal/adapter/handler/fiber/controller"
	"toktok-backend/internal/adapter/handler/fiber/router"
	"toktok-backend/internal/adapter/persistence/mysql"
	"toktok-backend/internal/adapter/persistence/mysql/repository"
	"toktok-backend/internal/core/service"
	"toktok-backend/internal/infra/config"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	ctx := context.Background()

	config, err := config.New("configs/.toml")
	if err != nil {
		panic(err)
	}

	client, err := mysql.NewClient(config)
	if err != nil {
		panic(err)
	}
	mysql.AutoMigration(ctx, client)

	userRepository := repository.NewUserRepository(client)
	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(userService)

	router := router.NewRouter(router.ControllerSet{
		UserController: userController,
	})

	if err := router.Listen(":8080"); err != nil {
		log.Fatal(err)
	}
}
