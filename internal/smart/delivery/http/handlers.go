package http

import (
	"github.com/gin-gonic/gin"
	"unraid-rest-api/internal/smart"
	"unraid-rest-api/service"
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
