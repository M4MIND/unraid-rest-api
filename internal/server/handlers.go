package server

import (
	cpuHttp "unraid-rest-api/internal/cpu/delivery/http"
	disksHttp "unraid-rest-api/internal/disks/delivery/http"
	dockerHttp "unraid-rest-api/internal/docker/delivery/http"
	memoryHttp "unraid-rest-api/internal/memory/delivery/http"
	networkHttp "unraid-rest-api/internal/network/delivery/http"
	raidHttp "unraid-rest-api/internal/raid/delivery/http"
	smartHttp "unraid-rest-api/internal/smart/delivery/http"
	"unraid-rest-api/service"
)

func (s *Server) MapHandlers(serviceContainer service.Container) {
	v1 := s.gin.Group("/api/v1")

	cpuHttp.MapRoutes(v1.Group("/cpu"), cpuHttp.NewHandler(serviceContainer))
	memoryHttp.MapRoutes(v1.Group("/memory"), memoryHttp.NewHandler(serviceContainer))
	networkHttp.MapRoutes(v1.Group("/network"), networkHttp.NewHandler(serviceContainer))
	dockerHttp.MapRoutes(v1.Group("/docker"), dockerHttp.NewHandler(serviceContainer))
	disksHttp.MapRoutes(v1.Group("/disks"), disksHttp.NewHandler(serviceContainer))
	raidHttp.MapRoutes(v1.Group("/raid"), raidHttp.NewHandler(serviceContainer))
	smartHttp.MapRoutes(v1.Group("/smart"), smartHttp.NewHandler(serviceContainer))
}
