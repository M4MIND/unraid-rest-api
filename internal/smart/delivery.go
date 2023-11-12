package smart

import "github.com/gin-gonic/gin"

type Handlers interface {
	GetSmartInfo() gin.HandlerFunc
}
