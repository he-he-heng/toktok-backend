package router

import "toktok-backend/internal/adapter/handler/fiber/controller"

type ControllerSet struct {
	UserController     *controller.UserController
	AvatarController   *controller.AvatarController
	RelationController *controller.RelationController
	RoomController     *controller.RoomController

	AuthController *controller.AuthController
}
