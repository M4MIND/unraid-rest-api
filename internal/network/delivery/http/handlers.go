package http

import (
	"github.com/gin-gonic/gin"
	"unraid-rest-api/internal/network"
	"unraid-rest-api/service"
)

type handler struct {
	services service.Container
}

func (h handler) GetAvgHistoryTick() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.JSON(200, h.services.NetworkService.GetLastHistory())
	}
}

func (h handler) GetAvgHistory() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.JSON(200, h.services.NetworkService.GetHistory())
	}
}

func NewHandler(s service.Container) network.Handlers {
	return &handler{services: s}
}
