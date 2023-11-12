package http

import (
	"github.com/gin-gonic/gin"
	"unraid-rest-api/internal/cpu"
	"unraid-rest-api/service"
)

type cpuHandler struct {
	services service.Container
}

func NewHandler(services service.Container) cpu.Handlers {
	return &cpuHandler{services: services}
}

func (s *cpuHandler) GetHistoryTick() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.JSON(200, s.services.CpuService.GetAvgHistoryLast())
	}
}

func (s *cpuHandler) GetHistory() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(200, s.services.CpuService.GetAvgHistory())
	}
}
func (s *cpuHandler) GetInfo() gin.HandlerFunc {
	return func(context *gin.Context) {
		info, err := s.services.CpuService.GetCpuInfo()

		if err != nil {
			context.JSON(500, gin.H{
				"error": err,
			})
		}

		context.JSON(200, info)
	}
}
