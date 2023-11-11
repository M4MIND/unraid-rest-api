package service

import (
	"unraid-rest-api/service/cpu"
	"unraid-rest-api/service/disks"
	"unraid-rest-api/service/docker"
	"unraid-rest-api/service/gpu"
	"unraid-rest-api/service/memory"
	"unraid-rest-api/service/network"
	"unraid-rest-api/service/raid"
	"unraid-rest-api/service/unraid"
)

type ServiceContainer struct {
	CpuService     *cpu.Service
	MemoryService  *memory.Service
	NetworkService *network.Service
	DockerService  *docker.Service
	DiskService    *disks.Service
	GpuService     *gpu.Service
	RaidService    *raid.Service
	UnraidService  *unraid.Service
}

func NewServiceContainer(
	cpuService *cpu.Service,
	memoryService *memory.Service,
	networkService *network.Service,
	dockerService *docker.Service,
	diskService *disks.Service,
	gpuService *gpu.Service,
	raidService *raid.Service,
	unraidService *unraid.Service,
) ServiceContainer {

	return ServiceContainer{
		CpuService:     cpuService,
		MemoryService:  memoryService,
		NetworkService: networkService,
		DockerService:  dockerService,
		DiskService:    diskService,
		GpuService:     gpuService,
		RaidService:    raidService,
		UnraidService:  unraidService,
	}
}
