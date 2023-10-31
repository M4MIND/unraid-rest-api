package server

import (
	"github.com/gin-gonic/gin"
	"log"
	"unraid-rest-api/service"
)

type Server struct {
	gin *gin.Engine
}

func NewServer() *Server {
	return &Server{gin: gin.Default()}
}

func (s *Server) Run(cpuSysstats *service.CpuSysstats, memorySysstats *service.MemorySysstats, sysstats *service.NetworkSysstats, client *service.DockerService) {
	s.MapHandlers(cpuSysstats, memorySysstats, sysstats, client)

	err := s.gin.Run("0.0.0.0:8554")

	if err != nil {
		log.Fatal(err)
	}

}
