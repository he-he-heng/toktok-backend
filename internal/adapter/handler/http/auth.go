package http

import (
	"net/http"
	"toktok-backend/internal/adapter/handler/http/dto"
	"toktok-backend/internal/adapter/handler/http/util"
	"toktok-backend/internal/core/port"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService port.AuthService
}

func NewAuthHandler(authService port.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

func (h *AuthHandler) Login(ctx *gin.Context) {
	var req dto.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		util.HandleErr(ctx, err)
		return
	}

	accessToken, refreshToken, err := h.authService.Login(ctx, req.UID, req.Password)
	if err != nil {
		util.HandleErr(ctx, err)
		return
	}

	util.HandleSuccess(ctx, http.StatusOK, "success create tokens", &dto.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})
}

func (h *AuthHandler) ReissueToken(ctx *gin.Context) {
	var req dto.ReissueTokenReqeust
	if err := ctx.ShouldBindJSON(&req); err != nil {
		util.HandleErr(ctx, err)
		return
	}

	accessToken, err := h.authService.ReissueToken(ctx, req.RefreshToken)
	if err != nil {
		util.HandleErr(ctx, err)
		return
	}

	util.HandleSuccess(ctx, http.StatusOK, "success reissue access token", &dto.ReissueTokenResponsea{
		AccessTokens: accessToken,
	})
}
