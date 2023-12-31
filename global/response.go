package global

import (
	"time"

	"github.com/gin-gonic/gin"
)

type commonResponse struct {
	State   bool   `json:"state"`
	Message string `json:"message"`
	Data    any    `json:"data"`
	Time    int64  `json:"time"`
}

func ReturnMessage(ctx *gin.Context, state bool, message string) {
	ctx.JSON(200, commonResponse{
		State:   state,
		Message: message,
		Time:    time.Now().Unix(),
	})
	ctx.Abort()
}

func ReturnData(ctx *gin.Context, state bool, message string, obj any) {
	ctx.JSON(200, commonResponse{
		State:   state,
		Message: message,
		Data:    obj,
		Time:    time.Now().Unix(),
	})
	ctx.Abort()
}
