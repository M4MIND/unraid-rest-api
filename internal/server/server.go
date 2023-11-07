package server

import (
	"log"
	"unraid-rest-api/service"

	"github.com/gin-gonic/gin"
)

type Server struct {
	gin *gin.Engine
}

func NewServer() *Server {
	return &Server{gin: gin.Default()}
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, Set-Cookie")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func (s *Server) Run(cpuSysstats *service.CpuSysstats, memorySysstats *service.MemorySysstats, networkSysstats *service.NetworkSysstats, dockerService *service.DockerService, sysstats *service.DisksSysstats) {
	s.gin.Use(corsMiddleware())

	s.MapHandlers(cpuSysstats, memorySysstats, networkSysstats, dockerService, sysstats)

	err := s.gin.Run("0.0.0.0:8554")

	if err != nil {
		log.Fatal(err)
	}

}
