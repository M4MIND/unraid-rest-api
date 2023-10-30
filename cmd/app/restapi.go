package app

import (
	"unraid-rest-api/controllers"
	"unraid-rest-api/services"

	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
)

var CpuController = new(controllers.CpuController)
var MemoryController = new(controllers.MemoryController)
var DockerController = new(controllers.DockerController)
var NetworkController = new(controllers.NetworkController)

var CpuService = new(services.CpuService)
var MemoryService = new(services.MemoryService)
var DockerService = new(services.DockerService)
var NetworkService = new(services.NetworkService)

func RestApi() {
	r := gin.Default()

	Gorutines()

	r.Use(cors.Default())

	api := r.Group("/api")
	{
		api.GET("/system/info/cpu", func(context *gin.Context) {
			CpuController.GetCpuInfo(context, CpuService)
		})
		api.GET("/system/info/cpu/stat", func(context *gin.Context) {
			CpuController.GetCpuStat(context, CpuService)
		})
		api.GET("/system/info/memory", func(context *gin.Context) {
			MemoryController.GetMemoryInfo(context, MemoryService)
		})

		api.GET("/system/info/memory/history", func(ctx *gin.Context) {
			MemoryController.GetMemoryHistory(ctx, MemoryService)
		})
		api.GET("/system/info/network", NetworkController.GetNetworks)
		api.GET("/system/info/network/history", func(context *gin.Context) {
			NetworkController.GetNetworksAvg(context, NetworkService)
		})

		api.GET("/system/info/docker/containers", func(context *gin.Context) {
			DockerController.GetContainerList(context, DockerService)
		})
		api.GET("/docker/container/command/:containerId/stop", func(context *gin.Context) {
			DockerController.StopContainerById(context, DockerService)
		})
		api.GET("/docker/container/command/:containerId/start", func(context *gin.Context) {
			DockerController.StartContainerById(context, DockerService)
		})
	}

	err := r.Run("0.0.0.0:8554")
	if err != nil {
		return
	} // listen and serve on 0.0.0.0:8080
}

func Gorutines() {
	go CpuService.Go()
	go MemoryService.Go()
	go NetworkService.Go()
}
