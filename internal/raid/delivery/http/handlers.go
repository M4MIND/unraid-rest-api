package http

import (
	"github.com/gin-gonic/gin"
	"unraid-rest-api/internal/raid"
	raidService "unraid-rest-api/service/raid"
)

type handler struct {
	raidService raidService.RaidService
}

func NewHandler(raidService raidService.RaidService) raid.Handlers {
	return handler{raidService: raidService}
}

func (h handler) GetMdstat() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.JSON(200, h.raidService.GetMdcmdStat())
	}
}
