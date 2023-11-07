package http

import (
	"github.com/gin-gonic/gin"
	"unraid-rest-api/internal/network"
)

func MapRoutes(group *gin.RouterGroup, h network.Handlers) {
	group.GET("/history", h.GetAvgHistory())
	group.GET("/history/tick", h.GetAvgHistoryTick())
}
