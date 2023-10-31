package main

import (
	"unraid-rest-api/internal/server"
	"unraid-rest-api/service"
)

func main() {
	https := server.NewServer()

	https.Run(
		service.CpuNew(),
		service.NewMemorySysstats(),
		service.NewNetworkSysstats(),
		service.NewDockerClient(),
	)
}
