package main

import (
	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
	"unraid-rest-api/controllers"
	"unraid-rest-api/services"
)

var CpuController = new(controllers.CpuController)
var MemoryController = new(controllers.MemoryController)
var DockerController = new(controllers.DockerController)
var NetworkController = new(controllers.NetworkController)

var CpuService = new(services.CpuService)
var MemoryService = new(services.MemoryService)

func main() {
	r := gin.Default()

	Gorutines()

	r.Use(cors.Default())
	r.GET("/system/info/cpu", func(context *gin.Context) {
		CpuController.GetCpuInfo(context, CpuService)
	})
	r.GET("/system/info/cpu/stat", func(context *gin.Context) {
		CpuController.GetCpuStat(context, CpuService)
	})
	r.GET("/system/info/memory", func(context *gin.Context) {
		MemoryController.GetMemoryInfo(context, MemoryService)
	})
	r.GET("/system/info/docker/containers", DockerController.GetContainerList)
	r.GET("/system/info/network", NetworkController.GetNetworks)

	err := r.Run("0.0.0.0:8554")
	if err != nil {
		return
	} // listen and serve on 0.0.0.0:8080
}

func Gorutines() {
	go CpuService.Go()
}
