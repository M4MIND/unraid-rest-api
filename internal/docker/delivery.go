package docker

import "github.com/gin-gonic/gin"

type Handlers interface {
	GetAllContainers() gin.HandlerFunc
}
