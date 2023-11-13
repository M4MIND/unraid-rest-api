package http

import (
	"unraid-rest-api/api/raid"

	"github.com/gin-gonic/gin"
)

func MapRoutes(g *gin.RouterGroup, h raid.Handlers) {
	g.GET("/mdcmdstat", h.GetMdstat())
}
