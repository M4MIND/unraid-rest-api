package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/rafacas/sysstats"
	"unraid-rest-api/services"
)

type NetworkController struct {
}

var prevNetworkAvg, _ = sysstats.GetNetRawStats()

func (ctrl *NetworkController) GetNetworks(c *gin.Context) {
	c.JSON(200, gin.H{
		"test": "ASD",
	})
}

func (ctrl *NetworkController) GetNetworksAvg(c *gin.Context, service *services.NetworkService) {
	c.JSON(200, service.GetHistory())
}
