package http

import (
	"github.com/gin-gonic/gin"
	"unraid-rest-api/internal/docker"
)

func MapRoutes(group *gin.RouterGroup, h docker.Handlers) {
	group.GET("/containers", h.GetAllContainers())
}
