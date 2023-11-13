package service

import (
	"unraid-rest-api/pkg/service/cpu"
	"unraid-rest-api/pkg/service/disks"
	"unraid-rest-api/pkg/service/docker"
	"unraid-rest-api/pkg/service/gpu"
	"unraid-rest-api/pkg/service/memory"
	"unraid-rest-api/pkg/service/network"
	"unraid-rest-api/pkg/service/raid"
	"unraid-rest-api/pkg/service/smart"
	"unraid-rest-api/pkg/service/unraid"
)

type Container struct {
	CpuService     *cpu.Service
	MemoryService  *memory.Service
	NetworkService *network.Service
	DockerService  *docker.Service
	DiskService    *disks.Service
	GpuService     *gpu.Service
	RaidService    *raid.Service
	UnraidService  *unraid.Service
	SmartService   *smart.Service
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
	smart *smart.Service,
) Container {

	container := Container{
		CpuService:     cpuService,
		MemoryService:  memoryService,
		NetworkService: networkService,
		DockerService:  dockerService,
		DiskService:    diskService,
		GpuService:     gpuService,
		RaidService:    raidService,
		UnraidService:  unraidService,
		SmartService:   smart,
	}

	return container
}
