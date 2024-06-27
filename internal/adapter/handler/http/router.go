package http

import (
	"toktok-backend/internal/adapter/config"
	"toktok-backend/internal/core/port"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type Router struct {
	*gin.Engine
}

func NewRouter(config *config.Config, token port.TokenService, userHandler UserHandler, authHandler AuthHandler) (*Router, error) {
	if config.Handler.Mode == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()
	router.Use(gin.Recovery())

	// custom validation
	v, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		err := v.RegisterValidation("user_role", userRoleValidator)
		if err != nil {
			return nil, err
		}

		err = v.RegisterValidation("regexp", regexpValidator)
		if err != nil {
			return nil, err
		}
	}

	api := router.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			users := v1.Group("/users")
			{
				users.POST("/", userHandler.Register)
			}

			auth := v1.Group("/auth")
			{
				auth.POST("/login", authHandler.Login)
				auth.POST("/reissue", authHandler.ReissueToken)
			}
		}
	}

	return &Router{
		router,
	}, nil
}

func (r *Router) Listen(addr string) error {
	return r.Run(addr)
}
