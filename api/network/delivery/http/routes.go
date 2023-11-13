package http

import (
	"unraid-rest-api/api/network"

	"github.com/gin-gonic/gin"
)

func MapRoutes(group *gin.RouterGroup, h network.Handlers) {
	group.GET("/history", h.GetAvgHistory())
	group.GET("/history/tick", h.GetAvgHistoryTick())
}
