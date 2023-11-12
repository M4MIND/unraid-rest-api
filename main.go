package main

import (
	"unraid-rest-api/internal/server"
	"unraid-rest-api/service"
	"unraid-rest-api/service/cpu"
	"unraid-rest-api/service/disks"
	"unraid-rest-api/service/docker"
	"unraid-rest-api/service/gpu"
	"unraid-rest-api/service/memory"
	"unraid-rest-api/service/network"
	"unraid-rest-api/service/raid"
	"unraid-rest-api/service/smart"
	"unraid-rest-api/service/unraid"
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

	http.Run(
		serviceContainer,
	)
}
