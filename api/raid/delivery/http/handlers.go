package http

import (
	"unraid-rest-api/api/raid"
	"unraid-rest-api/pkg/service"

	"github.com/gin-gonic/gin"
)

type handler struct {
	services service.Container
}

func NewHandler(s service.Container) raid.Handlers {
	return handler{services: s}
}

func (h handler) GetMdstat() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.JSON(200, h.services.RaidService.GetMdcmdStat())
	}
}
