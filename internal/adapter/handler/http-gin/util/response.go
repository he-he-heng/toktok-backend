package util

import (
	"toktok-backend/internal/adapter/handler/http/dto"

	"github.com/gin-gonic/gin"
)

// handleError determines the status code of an error and returns a JSON response with the error message and status code
func HandleErr(ctx *gin.Context, err error) {
	s := status(err)
	ctx.JSON(s, &dto.GeneralResponse{
		Status:   s,
		Messsage: parseErr(err),
		Data:     nil,
	})
}

func HandleSuccess(ctx *gin.Context, s int, m string, d ...any) {
	var data any = nil
	if len(d) > 0 {
		data = d[0]
	}

	ctx.JSON(s, &dto.GeneralResponse{
		Status:   s,
		Messsage: m,
		Data:     data,
	})
}
