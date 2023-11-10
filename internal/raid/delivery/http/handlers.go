package http

import (
	"github.com/gin-gonic/gin"
	"unraid-rest-api/internal/raid"
	raid2 "unraid-rest-api/service/raid"
)

type handler struct {
	raidService raid2.RaidService
}

func NewHandler(raidService raid2.RaidService) raid.Handlers {
	return handler{raidService: raidService}
}

func (h handler) GetMdstat() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.JSON(200, h.raidService.GetMdcmdStat())
	}
}
