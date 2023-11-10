package raid

import "github.com/gin-gonic/gin"

type Handlers interface {
	GetMdstat() gin.HandlerFunc
}
