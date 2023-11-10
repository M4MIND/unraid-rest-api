package server

import (
	cpuHttp "unraid-rest-api/internal/cpu/delivery/http"
	disksHttp "unraid-rest-api/internal/disks/delivery/http"
	dockerHttp "unraid-rest-api/internal/docker/delivery/http"
	memoryHttp "unraid-rest-api/internal/memory/delivery/http"
	networkHttp "unraid-rest-api/internal/network/delivery/http"
	raidHttp "unraid-rest-api/internal/raid/delivery/http"
	"unraid-rest-api/service/cpu"
	"unraid-rest-api/service/disks"
	"unraid-rest-api/service/docker"
	"unraid-rest-api/service/gpu"
	"unraid-rest-api/service/memory"
	"unraid-rest-api/service/network"
	"unraid-rest-api/service/raid"
)

func (s *Server) MapHandlers(cpuSysstats *cpu.CpuSysstats, memorySysstats *memory.MemorySysstats, networkSysstats *network.NetworkSysstats, client *docker.DockerService, sysstats *disks.DisksSysstats, gpuService *gpu.GpuService, raidService raid.RaidService) {
	v1 := s.gin.Group("/api/v1")

	cpuHttp.MapRoutes(v1.Group("/cpu"), cpuHttp.NewHandler(cpuSysstats))
	memoryHttp.MapRoutes(v1.Group("/memory"), memoryHttp.NewHandler(memorySysstats))
	networkHttp.MapRoutes(v1.Group("/network"), networkHttp.NewHandler(networkSysstats))
	dockerHttp.MapRoutes(v1.Group("/docker"), dockerHttp.NewHandler(client))
	disksHttp.MapRoutes(v1.Group("/disks"), disksHttp.NewHandler(sysstats))
	raidHttp.MapRoutes(v1.Group("/raid"), raidHttp.NewHandler(raidService))
}
