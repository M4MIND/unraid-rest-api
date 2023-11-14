package ws

import (
	"github.com/gin-gonic/gin"
	"time"
	"unraid-rest-api/api/websocket"
)

func MapRoutes(g *gin.RouterGroup, w websocket.Handlers) {
	server := NewWebsocket()
	///server.CreateTopic("ping-pong", w.PingPong, 10)
	server.CreateTopic("cpu-data", w.CpuState, time.Second)

	g.GET("", server.Handler())
}
