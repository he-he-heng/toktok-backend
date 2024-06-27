package http

import (
	"net/http"
	"toktok-backend/internal/adapter/handler/http/dto"
	"toktok-backend/internal/adapter/handler/http/util"
	"toktok-backend/internal/core/domain"
	"toktok-backend/internal/core/port"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService port.UserService
}

func NewUserHandler(userSerivce port.UserService) *UserHandler {
	return &UserHandler{
		userService: userSerivce,
	}
}

func (h *UserHandler) Register(ctx *gin.Context) {
	var req dto.RegisterReqeust
	if err := ctx.ShouldBindJSON(&req); err != nil {
		util.HandleErr(ctx, err)
		return
	}

	user := domain.User{
		UID:      req.UID,
		Password: req.Password,
	}

	_, err := h.userService.Register(ctx, &user)
	if err != nil {
		util.HandleErr(ctx, err)
		return
	}

	util.HandleSuccess(ctx, http.StatusOK, "success register")
}
