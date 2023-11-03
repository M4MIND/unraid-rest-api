package http

import (
	"github.com/gin-gonic/gin"
	"unraid-rest-api/internal/disks"
)

func MapRoutes(group *gin.RouterGroup, handlers disks.Handlers) {
	group.GET("/history", handlers.GetHistory())
}