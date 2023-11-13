package http

import (
	"unraid-rest-api/api/docker"

	"github.com/gin-gonic/gin"
)

func MapRoutes(group *gin.RouterGroup, h docker.Handlers) {
	group.GET("/containers", h.GetAllContainers())
}
