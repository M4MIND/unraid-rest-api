package http

import (
	"github.com/gin-gonic/gin"
	"unraid-rest-api/api/cpu"
)

type Test struct {
	Message string `json:"message"`
}

func MapRoutes(group *gin.RouterGroup, handler cpu.Handlers) {
	group.GET("/history", handler.GetHistory())
	group.GET("/history/tick", handler.GetHistoryTick())
	group.GET("/info", handler.GetInfo())
}
