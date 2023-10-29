package controllers

import (
	"github.com/gin-gonic/gin"
	"unraid-rest-api/services"
)

type DockerController struct {
}

func (ctrl *DockerController) GetContainerList(c *gin.Context, service *services.DockerService) {
	c.JSON(200, service.GetContainers())
}

func (ctrl *DockerController) StopContainerById(c *gin.Context, service *services.DockerService) {
	service.StopContainers(c.Param("containerId"))

	c.JSON(200, gin.H{
		"id": c.Param("containerId"),
	})
}

func (ctrl *DockerController) StartContainerById(c *gin.Context, service *services.DockerService) {
	service.StartContainers(c.Param("containerId"))
	c.JSON(200, gin.H{
		"id": c.Param("containerId"),
	})
}
