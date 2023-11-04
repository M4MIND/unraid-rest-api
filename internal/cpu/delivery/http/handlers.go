package http

import (
	"unraid-rest-api/internal/cpu"
	"unraid-rest-api/service"

	"github.com/gin-gonic/gin"
)

type cpuHandler struct {
	cpuService *service.CpuSysstats
}

func NewHandler(cpu *service.CpuSysstats) cpu.Handlers {
	return &cpuHandler{cpuService: cpu}
}

func (s *cpuHandler) GetHistoryTick() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.JSON(200, s.cpuService.GetAvgHistoryLast())
	}
}

func (s *cpuHandler) GetHistory() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(200, s.cpuService.GetAvgHistory())
	}
}
func (s *cpuHandler) GetInfo() gin.HandlerFunc {
	return func(context *gin.Context) {
		info, err := s.cpuService.GetCpuInfo()

		if err != nil {
			context.JSON(500, gin.H{
				"error": err,
			})
		}

		context.JSON(200, info)
	}
}
