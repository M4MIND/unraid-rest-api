package docker

import (
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"log"
)

type Service struct {
	client *client.Client
}

func NewService() *Service {
	client, _ := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	return &Service{
		client: client,
	}
}

func (c *Service) GetAllContainers() []types.Container {
	containers, err := c.client.ContainerList(context.Background(), types.ContainerListOptions{
		All: true,
	})

	if err != nil {
		log.Fatalln(err)
	}

	return containers
}
