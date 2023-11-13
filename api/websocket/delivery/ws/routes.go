package ws

import (
	"github.com/gin-gonic/gin"
	"unraid-rest-api/api/websocket"
)

func MapRoutes(g *gin.RouterGroup, w websocket.Handlers) {
	server := NewWebsocket()

	server.AddTopic("cpu-state", w.CpuState)

	g.GET("", server.Handler())
}
