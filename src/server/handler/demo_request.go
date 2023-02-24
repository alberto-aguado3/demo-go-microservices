package handler

import (
	"demo/src/service"

	"github.com/gin-gonic/gin"
)

// ignore this handler, just for testing http as demo
func DemoRequest(s service.AlbumService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		_ = s.GetWeb(ctx)
		ctx.AbortWithStatus(200)
	}
}
