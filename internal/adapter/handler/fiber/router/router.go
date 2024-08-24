package router

import (
	"toktok-backend/internal/adapter/handler/fiber/catcher"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type Router struct {
	controllerSet ControllerSet
	app           *fiber.App
}

func NewRouter(controllerSet ControllerSet) *Router {
	router := Router{
		controllerSet: controllerSet,
	}

	// new app
	config := fiber.Config{
		ErrorHandler: catcher.CustomErrorHandler,
	}
	app := fiber.New(config)

	// set middleware
	app.Use(logger.New())

	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	// set token setting

	// set my api
	v1 := app.Group("/v1")
	{
		api := v1.Group("/api")
		{
			users := api.Group("/users")
			{
				users.Get("/", router.controllerSet.UserController.ListUser)
				users.Get("/:id", router.controllerSet.UserController.GetUser)
				users.Post("/", router.controllerSet.UserController.CreateUser)
				users.Put("/:id", router.controllerSet.UserController.UpdateUser)
				users.Delete("/:id", router.controllerSet.UserController.DeleteUser)
			}
		}
	}

	router.app = app
	return &router
}

func (r *Router) Listen(address string) error {
	return r.app.Listen(address)
}
