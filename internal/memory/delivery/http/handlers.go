package http

import (
	"github.com/gin-gonic/gin"
	"unraid-rest-api/internal/memory"
	"unraid-rest-api/service"
)

type handler struct {
	services service.ServiceContainer
}

func NewHandler(s service.ServiceContainer) memory.Handlers {
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
