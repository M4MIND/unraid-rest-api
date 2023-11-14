package ws

import (
	"github.com/gin-gonic/gin"
	"unraid-rest-api/api/websocket"
)

func MapRoutes(g *gin.RouterGroup, w websocket.Handlers) {
	server := NewWebsocket()
	server.AddTopic("cpu-data", w.CpuState, 1)

	g.GET("", server.Handler())
}
