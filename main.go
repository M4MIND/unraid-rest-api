package main

import (
	"unraid-rest-api/api/server"
	disksMonitoring "unraid-rest-api/pkg/monitoring/disks"
	"unraid-rest-api/pkg/service"
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

func main() {
	http := server.NewServer()

	serviceContainer := service.NewServiceContainer(
		cpu.NewService(),
		memory.NewService(),
		network.NewService(),
		docker.NewService(),
		disks.NewService(),
		gpu.NewService(),
		raid.NewService(),
		unraid.NewService(),
		smart.NewService(),
	)

	go disksMonitoring.New(serviceContainer).Work()

	http.Run(serviceContainer)
}
