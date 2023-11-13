package http

import (
	"unraid-rest-api/internal/websocket"

	"github.com/gin-gonic/gin"
)

func MapRoutes(group *gin.RouterGroup, handler websocket.Handlers) {
	group.GET("/websocket", handler.UpgradeWebsocket())
}
