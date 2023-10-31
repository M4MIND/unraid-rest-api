package service

import (
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"log"
)

type DockerService struct {
	client *client.Client
}

func NewDockerClient() *DockerService {
	client, _ := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	return &DockerService{
		client: client,
	}
}

func (c *DockerService) GetAllContainers() []types.Container {
	containers, err := c.client.ContainerList(context.Background(), types.ContainerListOptions{
		All: true,
	})

	if err != nil {
		log.Fatalln(err)
	}

	return containers
}
