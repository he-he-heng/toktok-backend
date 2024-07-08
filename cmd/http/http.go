package main

import (
	"context"
	"fmt"
	"log/slog"
	"toktok-backend/internal/adapter/auth/token"
	"toktok-backend/internal/adapter/config"
	"toktok-backend/internal/adapter/handler/http"
	"toktok-backend/internal/adapter/handler/http/router"
	"toktok-backend/internal/adapter/logger"
	"toktok-backend/internal/adapter/persistence/mysql"
	"toktok-backend/internal/adapter/persistence/mysql/repository"
	"toktok-backend/internal/core/service"
)

func main() {
	config, err := config.New(".toml")
	if err != nil {
		panic(err)
	}

	logger.Set(config)

	slog.Info("Starting the application")

	db, err := mysql.New(config)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	slog.Info("Successfully connected to the database")

	ctx := context.Background()
	err = db.AutoMigration(ctx)
	if err != nil {
		panic(err)
	}

	slog.Info("Successfully migrate the database")

	tokenService, err := token.New(config)
	if err != nil {
		panic(err)
	}

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userHandler := http.NewUserHandler(userService)

	authService := service.NewAuthService(userRepository, tokenService)
	authHandler := http.NewAuthHandler(authService)

	router := router.New(userHandler, authHandler, tokenService)
	addr := fmt.Sprintf(":%s", config.Handler.Port)

	slog.Info("Starting the HTTP server", "listen_address", addr)
	err = router.Listen(addr)
	if err != nil {
		slog.Error("Error starting the HTTP server", "error", err)
	}
}
