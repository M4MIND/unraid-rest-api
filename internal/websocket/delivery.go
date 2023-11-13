package websocket

import "github.com/gin-gonic/gin"

type Handlers interface {
	UpgradeWebsocket() gin.HandlerFunc
}
