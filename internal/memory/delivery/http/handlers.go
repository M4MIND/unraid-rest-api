package http

import (
	"github.com/gin-gonic/gin"
	"unraid-rest-api/internal/memory"
	"unraid-rest-api/service"
)

type handler struct {
	memoryService *service.MemorySysstats
}

func NewHandler(memorySysstats *service.MemorySysstats) memory.Handlers {
	return &handler{memoryService: memorySysstats}
}

func (h handler) GetHistory() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.JSON(200, h.memoryService.GetHistory())
	}
}
