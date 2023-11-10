package http

import (
	"github.com/gin-gonic/gin"
	"unraid-rest-api/internal/disks"
	disks2 "unraid-rest-api/service/disks"
)

type disksHandler struct {
	disksService *disks2.DisksSysstats
}

func (d disksHandler) GetDisksInfoLsblk() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.JSON(200, d.disksService.GetDisksLsblk())
	}
}

func (d disksHandler) GetHistoryTick() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.JSON(200, d.disksService.GetHistoryLast())
	}
}

func (d disksHandler) GetHistory() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.JSON(200, d.disksService.GetHistory())
	}
}

func NewHandler(s *disks2.DisksSysstats) disks.Handlers {
	return &disksHandler{
		disksService: s,
	}
}
