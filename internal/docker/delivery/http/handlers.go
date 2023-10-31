package http

import (
	"github.com/gin-gonic/gin"
	"unraid-rest-api/internal/docker"
	"unraid-rest-api/service"
)

type dockerHandler struct {
	dockerService *service.DockerService
}

func (d dockerHandler) GetAllContainers() gin.HandlerFunc {
	//TODO implement me
	panic("implement me")
}

func NewHandler(s *service.DockerService) docker.Handlers {
	return &dockerHandler{
		dockerService: s,
	}
}
