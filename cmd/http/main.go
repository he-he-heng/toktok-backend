package main

import (
	"context"
	"log"
	"toktok-backend/internal/adapter/auth/jwt"
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

	cfg, err := config.New("configs/.toml")
	if err != nil {
		panic(err)
	}

	client, err := mysql.NewClient(cfg)
	if err != nil {
		panic(err)
	}

	jwtService, err := jwt.NewJWT(cfg)
	if err != nil {
		panic(err)
	}

	err = mysql.AutoMigration(ctx, client)
	if err != nil {
		panic(err)
	}

	userRepository := repository.NewUserRepository(client)
	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(userService)

	authService := service.NewAuthService(userRepository, jwtService)
	authController := controller.NewAuthController(authService)

	avatarRepository := repository.NewAvatarRepository(client)
	avatarService := service.NewAvatarService(avatarRepository)
	avatarController := controller.NewAvatarController(avatarService)

	roomRepository := repository.NewRoomRepoistory(client)
	// roomService := service

	relationRepository := repository.NewRelationRepository(client)
	relationService := service.NewRelationService(relationRepository, roomRepository)
	relationController := controller.NewRelationController(relationService)

	r := router.NewRouter(
		router.ControllerSet{
			UserController:     userController,
			AvatarController:   avatarController,
			RelationController: relationController,

			AuthController: authController,
		},

		router.ServiceSet{
			JWTService: jwtService,
		},
	)

	if err := r.Listen(":8080"); err != nil {
		log.Fatal(err)
	}
}
