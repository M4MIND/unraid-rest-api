package main

import (
	"unraid-rest-api/internal/server"
	"unraid-rest-api/service/cpu"
	"unraid-rest-api/service/disks"
	"unraid-rest-api/service/docker"
	"unraid-rest-api/service/gpu"
	"unraid-rest-api/service/memory"
	"unraid-rest-api/service/network"
	"unraid-rest-api/service/raid"
)

func main() {
	http := server.NewServer()

	http.Run(
		cpu.NewCpuSysstats(),
		memory.NewMemorySysstats(),
		network.NewNetworkSysstats(),
		docker.NewDockerClient(),
		disks.NewDisksSysstats(),
		gpu.NewGpuService(),
		raid.NewRaidService(),
	)
}
