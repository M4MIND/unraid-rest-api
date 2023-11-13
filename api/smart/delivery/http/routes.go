package http

import (
	"unraid-rest-api/api/smart"

	"github.com/gin-gonic/gin"
)

func MapRoutes(g *gin.RouterGroup, h smart.Handlers) {
	g.GET("/info/:disk", h.GetSmartInfo())
}
