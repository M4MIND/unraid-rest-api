package http

import (
	"github.com/gin-gonic/gin"
	"unraid-rest-api/internal/smart"
)

func MapRoutes(g *gin.RouterGroup, h smart.Handlers) {
	g.GET("/info/:disk", h.GetSmartInfo())
}
