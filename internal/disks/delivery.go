package disks

import "github.com/gin-gonic/gin"

type Handlers interface {
	GetHistory() gin.HandlerFunc
}