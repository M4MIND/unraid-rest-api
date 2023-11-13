package http

import (
	"unraid-rest-api/api/smart"
	"unraid-rest-api/pkg/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	services service.Container
}

func (h *Handler) GetSmartInfo() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.JSON(200, h.services.SmartService.GetDiskSmartInfo("sdb"))
	}
}

func NewHandler(s service.Container) smart.Handlers {
	return &Handler{services: s}
}
