package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"unraid-rest-api/internal/disks"
	"unraid-rest-api/service"
)

type disksHandler struct {
	services service.Container
}

func (d disksHandler) GetArrayInfo() gin.HandlerFunc {
	return func(context *gin.Context) {
		mdstat := d.services.RaidService.GetMdcmdStat()

		fmt.Println(mdstat)

		context.JSON(200, mdstat)
	}
}

func (d disksHandler) GetDisksInfoLsblk() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.JSON(200, d.services.DiskService.GetDisksLsblk())
	}
}

func (d disksHandler) GetHistoryTick() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.JSON(200, d.services.DiskService.GetHistoryLast())
	}
}

func (d disksHandler) GetHistory() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.JSON(200, d.services.DiskService.GetHistory())
	}
}

func NewHandler(s service.Container) disks.Handlers {
	return &disksHandler{
		services: s,
	}
}
