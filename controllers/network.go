package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/rafacas/sysstats"
)

type NetworkController struct {
}

var prevNetworkAvg, _ = sysstats.GetNetRawStats()

func (ctrl *NetworkController) GetNetworks(c *gin.Context) {
	network, _ := sysstats.GetNetRawStats()

	avg, _ := sysstats.GetNetAvgStats(network, prevNetworkAvg)

	prevNetworkAvg = network
	c.JSON(200, gin.H{
		"test": avg,
	})
}
