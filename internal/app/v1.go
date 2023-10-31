package app

import (
	v1 "unraid-rest-api/internal/controller/http/api/v1"
	"unraid-rest-api/services"

	"github.com/gin-gonic/gin"
)

var CpuController = new(v1.CpuController)
var MemoryController = new(v1.MemoryController)
var DockerController = new(v1.DockerController)
var NetworkController = new(v1.NetworkController)

var CpuService = new(services.CpuService)
var MemoryService = new(services.MemoryService)
var DockerService = new(services.DockerService)
var NetworkService = new(services.NetworkService)

func RestApi() {
	handler := gin.New()

	Gorutines()

	v1.NewRouter(handler)

	err := handler.Run("0.0.0.0:8554")
	if err != nil {
		return
	} // listen and serve on 0.0.0.0:8080
}

func Gorutines() {
	go CpuService.Go()
	go MemoryService.Go()
	go NetworkService.Go()
}
