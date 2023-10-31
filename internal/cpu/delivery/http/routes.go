package http

import (
	"unraid-rest-api/internal/cpu"

	"github.com/gin-gonic/gin"
)

func MapRoutes(group *gin.RouterGroup, handler cpu.Handlers) {
	group.GET("/history", handler.GetHistory())
	group.GET("/info", handler.GetInfo())
}
