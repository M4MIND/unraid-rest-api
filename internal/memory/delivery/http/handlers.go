package http

import (
	"github.com/gin-gonic/gin"
	"unraid-rest-api/internal/memory"
	memory2 "unraid-rest-api/service/memory"
)

type handler struct {
	memoryService *memory2.MemorySysstats
}

func NewHandler(memorySysstats *memory2.MemorySysstats) memory.Handlers {
	return &handler{memoryService: memorySysstats}
}

func (h handler) GetHistory() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.JSON(200, h.memoryService.GetHistory())
	}
}

func (h handler) GetLastTick() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.JSON(200, h.memoryService.GetHistoryLast())
	}
}
