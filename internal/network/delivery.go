package network

import "github.com/gin-gonic/gin"

type Handlers interface {
	GetAvgHistory() gin.HandlerFunc
}
