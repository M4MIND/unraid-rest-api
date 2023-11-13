package http

import (
	"unraid-rest-api/api/network"
	"unraid-rest-api/pkg/service"

	"github.com/gin-gonic/gin"
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
