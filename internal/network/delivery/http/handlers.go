package http

import (
	"github.com/gin-gonic/gin"
	"unraid-rest-api/internal/network"
	network2 "unraid-rest-api/service/network"
)

type handler struct {
	networkSysstats *network2.NetworkSysstats
}

func (h handler) GetAvgHistoryTick() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.JSON(200, h.networkSysstats.GetLastHistory())
	}
}

func (h handler) GetAvgHistory() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.JSON(200, h.networkSysstats.GetHistory())
	}
}

func NewHandler(n *network2.NetworkSysstats) network.Handlers {
	return &handler{networkSysstats: n}
}
