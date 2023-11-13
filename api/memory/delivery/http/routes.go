package http

import (
	"unraid-rest-api/api/memory"

	"github.com/gin-gonic/gin"
)

func MapRoutes(group *gin.RouterGroup, handler memory.Handlers) {
	group.GET("/history", handler.GetHistory())
	group.GET("/tick", handler.GetLastTick())
}
