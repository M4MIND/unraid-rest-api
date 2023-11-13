package http

import (
	"unraid-rest-api/api/disks"

	"github.com/gin-gonic/gin"
)

func MapRoutes(group *gin.RouterGroup, handlers disks.Handlers) {
	group.GET("/history", handlers.GetHistory())
	group.GET("/history/tick", handlers.GetHistoryTick())
	group.GET("/info/lsblk", handlers.GetDisksInfoLsblk())
	group.GET("/array/info", handlers.GetArrayInfo())
}
