package http

import (
	"github.com/gin-gonic/gin"
	"unraid-rest-api/internal/docker"
	"unraid-rest-api/service"
)

type dockerHandler struct {
	dockerService *service.DockerService
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
		context.JSON(200, d.dockerService.GetAllContainers())
	}
}

func NewHandler(s *service.DockerService) docker.Handlers {
	return &dockerHandler{
		dockerService: s,
	}
}
