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

// TODO: add arg config
func NewRouter(controllerSet ControllerSet) *Router {
	router := Router{
		controllerSet: controllerSet,
	}

	// new app
	config := fiber.Config{
		ErrorHandler: catcher.CustomErrorHandler,
		// EnablePrintRoutes: true,
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

			avatars := api.Group("/avatars")
			{
				avatars.Get("/", router.controllerSet.AvatarController.ListAvatar)
				avatars.Get("/:id", router.controllerSet.AvatarController.GetAvatar)
				avatars.Post("/", router.controllerSet.AvatarController.CreateAvatar)
				avatars.Put("/:id", router.controllerSet.AvatarController.UpdateAvatar)
				avatars.Delete("/:id", router.controllerSet.AvatarController.DeleteAvatar)
			}

			relations := api.Group("/relations")
			{
				relations.Get("/", router.controllerSet.RelationController.ListRelation)
				relations.Get("/:id", router.controllerSet.RelationController.GetRelation)
				relations.Post("/", router.controllerSet.RelationController.CreateRelation)
				relations.Put("/:id", router.controllerSet.RelationController.UpdateRelation)
				relations.Delete("/:id", router.controllerSet.RelationController.DeleteRelation)
			}
		}
	}

	router.app = app
	return &router
}

func (r *Router) Listen(address string) error {
	return r.app.Listen(address)
}
