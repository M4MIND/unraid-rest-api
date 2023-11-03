package http

import (
	"github.com/gin-gonic/gin"
	"unraid-rest-api/internal/memory"
)

func MapRoutes(group *gin.RouterGroup, handler memory.Handlers) {
	group.GET("/history", handler.GetHistory())
	group.GET("/tick", handler.GetLastTick())
}
