package http

import (
	"github.com/gin-gonic/gin"
	"unraid-rest-api/internal/disks"
	"unraid-rest-api/service"
)

type disksHandler struct {
	disksService *service.DisksSysstats
}

func (d disksHandler) GetHistory() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.JSON(200, d.disksService.GetHistory())
	}
}

func NewHandler(s *service.DisksSysstats) disks.Handlers {
	return &disksHandler{
		disksService: s,
	}
}
