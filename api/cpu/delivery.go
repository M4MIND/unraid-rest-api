package cpu

import "github.com/gin-gonic/gin"

type Handlers interface {
	GetHistory() gin.HandlerFunc
	GetHistoryTick() gin.HandlerFunc
	GetInfo() gin.HandlerFunc
}
