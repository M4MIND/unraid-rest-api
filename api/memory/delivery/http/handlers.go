package http

import (
	"unraid-rest-api/api/memory"
	"unraid-rest-api/pkg/service"

	"github.com/gin-gonic/gin"
)

type handler struct {
	services service.Container
}

func NewHandler(s service.Container) memory.Handlers {
	return &handler{services: s}
}

func (h handler) GetHistory() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.JSON(200, h.services.MemoryService.GetHistory())
	}
}

func (h handler) GetLastTick() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.JSON(200, h.services.MemoryService.GetHistoryLast())
	}
}
