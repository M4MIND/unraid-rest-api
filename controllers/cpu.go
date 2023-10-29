package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/rafacas/sysstats"
	"unraid-rest-api/services"
)

type CpuController struct {
}

var tempCpuStats, err = sysstats.GetCpuRawStats()

func (ctrl *CpuController) GetCpuInfo(c *gin.Context, cpuService *services.CpuService) {
	c.JSON(200, gin.H{
		"avg": cpuService.GetLastCpuAvg(),
	})
}

func (ctrl *CpuController) GetCpuStat(c *gin.Context, cpuService *services.CpuService) {
	c.JSON(200, cpuService.GetAvgHistory())
}
