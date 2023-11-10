package http

import (
	"github.com/gin-gonic/gin"
	"unraid-rest-api/internal/raid"
)

func MapRoutes(g *gin.RouterGroup, h raid.Handlers) {
	g.GET("/mdcmdstat", h.GetMdstat())
}
