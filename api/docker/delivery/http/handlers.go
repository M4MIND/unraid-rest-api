package http

import (
	"unraid-rest-api/api/docker"
	"unraid-rest-api/pkg/service"

	"github.com/gin-gonic/gin"
)

type dockerHandler struct {
	services service.Container
}

func (d dockerHandler) StopContainer(id string) gin.HandlerFunc {
	//TODO implement me
	panic("implement me")
}

func (d dockerHandler) StartContainer(id string) gin.HandlerFunc {
	//TODO implement me
	panic("implement me")
}

func (d dockerHandler) RestartContainer(id string) gin.HandlerFunc {
	//TODO implement me
	panic("implement me")
}

func (d dockerHandler) KillContainer(id string) gin.HandlerFunc {
	//TODO implement me
	panic("implement me")
}

func (d dockerHandler) GetAllContainers() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.JSON(200, d.services.DockerService.GetAllContainers())
	}
}

func NewHandler(s service.Container) docker.Handlers {
	return &dockerHandler{
		services: s,
	}
}
