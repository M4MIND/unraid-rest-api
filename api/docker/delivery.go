package docker

import "github.com/gin-gonic/gin"

type Handlers interface {
	GetAllContainers() gin.HandlerFunc
	StopContainer(id string) gin.HandlerFunc
	StartContainer(id string) gin.HandlerFunc
	RestartContainer(id string) gin.HandlerFunc
	KillContainer(id string) gin.HandlerFunc
}
