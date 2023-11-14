package server

import (
	cpuHttp "unraid-rest-api/api/cpu/delivery/http"
	disksHttp "unraid-rest-api/api/disks/delivery/http"
	dockerHttp "unraid-rest-api/api/docker/delivery/http"
	memoryHttp "unraid-rest-api/api/memory/delivery/http"
	networkHttp "unraid-rest-api/api/network/delivery/http"
	raidHttp "unraid-rest-api/api/raid/delivery/http"
	smartHttp "unraid-rest-api/api/smart/delivery/http"
	websocketWs "unraid-rest-api/api/websocket/delivery/ws"
	"unraid-rest-api/pkg/service"
)

func (s *Server) MapHandlers(serviceContainer service.Container) {
	v1 := s.Gin.Group("/api/v1")

	cpuHttp.MapRoutes(v1.Group("/cpu"), cpuHttp.NewHandler(serviceContainer))
	memoryHttp.MapRoutes(v1.Group("/memory"), memoryHttp.NewHandler(serviceContainer))
	networkHttp.MapRoutes(v1.Group("/network"), networkHttp.NewHandler(serviceContainer))
	dockerHttp.MapRoutes(v1.Group("/docker"), dockerHttp.NewHandler(serviceContainer))
	disksHttp.MapRoutes(v1.Group("/disks"), disksHttp.NewHandler(serviceContainer))
	raidHttp.MapRoutes(v1.Group("/raid"), raidHttp.NewHandler(serviceContainer))
	smartHttp.MapRoutes(v1.Group("/smart"), smartHttp.NewHandler(serviceContainer))
	websocketWs.MapRoutes(v1.Group("/ws"), websocketWs.NewHandler(serviceContainer))
}
