package server

import (
	"log"
	"unraid-rest-api/service/cpu"
	"unraid-rest-api/service/disks"
	"unraid-rest-api/service/docker"
	"unraid-rest-api/service/gpu"
	"unraid-rest-api/service/memory"
	"unraid-rest-api/service/network"
	"unraid-rest-api/service/raid"

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

func (s *Server) Run(
	cpuSysstats *cpu.CpuSysstats,
	memorySysstats *memory.MemorySysstats,
	networkSysstats *network.NetworkSysstats,
	dockerService *docker.DockerService,
	sysstats *disks.DisksSysstats,
	gpuService *gpu.GpuService,
	raidService raid.RaidService) {
	s.gin.Use(corsMiddleware())

	s.MapHandlers(cpuSysstats, memorySysstats, networkSysstats, dockerService, sysstats, gpuService, raidService)

	err := s.gin.Run("0.0.0.0:8554")

	if err != nil {
		log.Fatal(err)
	}

}
