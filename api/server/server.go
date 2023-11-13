package server

import (
	"log"
	"unraid-rest-api/pkg/service"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Gin *gin.Engine
}

func NewServer() *Server {
	return &Server{Gin: gin.Default()}
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

func (s *Server) Run(serviceContainer service.Container) {
	s.Gin.Use(corsMiddleware())
	s.MapHandlers(serviceContainer)

	err := s.Gin.Run("0.0.0.0:8554")

	if err != nil {
		log.Fatal(err)
	}
}
