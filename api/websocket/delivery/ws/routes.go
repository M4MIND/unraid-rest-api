package ws

import (
	"github.com/gin-gonic/gin"
	"time"
	"unraid-rest-api/api/websocket"
)

func MapRoutes(g *gin.RouterGroup, w websocket.Handlers) {
	server := NewWebsocket()
	server.CreateTopic("ping-pong", w.PingPong, websocket.HandlerParams{SendOnce: false, Sleep: time.Second * 10})
	server.CreateTopic("cpu-data", w.CpuState, websocket.HandlerParams{SendOnce: false, Sleep: time.Second})
	server.CreateTopic("array-info", w.ArrayInfo, websocket.HandlerParams{SendOnce: false, Sleep: time.Second * 10})
	server.CreateTopic("memory-info-tick", w.MemoryInfoTick, websocket.HandlerParams{SendOnce: false, Sleep: time.Second})
	server.CreateTopic("stats-cpu-once", w.CpuStats, websocket.HandlerParams{SendOnce: true, Sleep: time.Second})
	server.CreateTopic("stats-cpu-tick", w.CpuStatsTick, websocket.HandlerParams{SendOnce: false, Sleep: time.Second})

	g.GET("/topics", server.Handler())
}
