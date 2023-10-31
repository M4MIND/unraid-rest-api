package server

import (
	cpuHttp "unraid-rest-api/internal/cpu/delivery/http"
	dockerHttp "unraid-rest-api/internal/docker/delivery/http"
	memoryHttp "unraid-rest-api/internal/memory/delivery/http"
	networkHttp "unraid-rest-api/internal/network/delivery/http"
	"unraid-rest-api/service"
)

func (s *Server) MapHandlers(cpuSysstats *service.CpuSysstats, memorySysstats *service.MemorySysstats, networkSysstats *service.NetworkSysstats, client *service.DockerService) {
	v1 := s.gin.Group("/api/v1")

	cpuHttp.MapRoutes(v1.Group("/cpu"), cpuHttp.NewHandler(cpuSysstats))
	memoryHttp.MapRoutes(v1.Group("/memory"), memoryHttp.NewHandler(memorySysstats))
	networkHttp.MapRoutes(v1.Group("/network"), networkHttp.NewHandler(networkSysstats))
	dockerHttp.MapRoutes(v1.Group("/docker"), dockerHttp.NewHandler(client))
}
