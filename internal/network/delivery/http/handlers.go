package http

import (
	"github.com/gin-gonic/gin"
	"unraid-rest-api/internal/network"
	"unraid-rest-api/service"
)

type handler struct {
	networkSysstats *service.NetworkSysstats
}

func (h handler) GetAvgHistory() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.JSON(200, h.networkSysstats.GetHistory())
	}
}

func NewHandler(n *service.NetworkSysstats) network.Handlers {
	return &handler{networkSysstats: n}
}
