package main

import (
	"fmt"
	"unraid-rest-api/internal/server"
	"unraid-rest-api/service"
)

func main() {
	http := server.NewServer()

	nvidia := service.NewGpuService()

	info, _ := nvidia.GetInfo()

	fmt.Println(info.Gpu.FbMemoryUsage.Total)

	http.Run(
		service.NewCpuSysstats(),
		service.NewMemorySysstats(),
		service.NewNetworkSysstats(),
		service.NewDockerClient(),
		service.NewDisksSysstats(),
	)
}
