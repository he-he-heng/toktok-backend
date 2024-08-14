package router

import (
	"toktok-backend/internal/adapter/handler/http"
	"toktok-backend/internal/adapter/handler/http/dto"
	"toktok-backend/internal/adapter/handler/http/middleware"
	"toktok-backend/internal/adapter/handler/http/myerror"
	"toktok-backend/internal/core/port"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type Router struct {
	App *fiber.App
}

func New(userHandler *http.UserHandler, authHandler *http.AuthHandler, tokenService port.TokenService) *Router {

	app := fiber.New(fiber.Config{
		ErrorHandler: myerror.Handler,
	})

	app.Use(logger.New())
	// Or extend your config for customization
	// Logging remote IP and Port
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	// app.Use(recover.New())
	// recover,

	m := middleware.New(tokenService)

	api := app.Group("/api")

	app.Post("/api/v1/auth/token-validation", m.ValidateToken, func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusOK)
	})

	{
		v1 := api.Group("v1")
		users := v1.Group("/users")
		{
			users.Post("/", userHandler.Register)
		}

		auth := v1.Group("/auth")
		{
			auth.Post("/login", authHandler.Login)

			token := auth.Group("/token")
			{
				token.Post("-refresh", authHandler.Refresh)

			}

		}
	}

	test := app.Group("/test", m.ValidateToken)
	{
		test.Post("/token/validation", func(c *fiber.Ctx) error {
			data := c.Locals(middleware.AuthorizationPayloadKey)
			return c.JSON(&dto.General{
				Status:  fiber.StatusOK,
				Message: "성공적으로 인증을 완료했습니다.",
				Data:    data,
			})
		})
	}

	return &Router{
		App: app,
	}
}

func (r *Router) Listen(addr string) error {
	return r.App.Listen(addr)
}
