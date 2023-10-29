package services

import (
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"log"
)

type DockerService struct {
}

var DockerClient, _ = client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())

func (ctrl *DockerService) GetContainers() []types.Container {
	containers, err := DockerClient.ContainerList(context.Background(), types.ContainerListOptions{
		All: true,
	})
	if err != nil {
		panic(err)
	}

	return containers
}

func (ctrl *DockerService) StopContainers(containerId string) {
	err := DockerClient.ContainerStop(context.Background(), containerId, container.StopOptions{})

	if err != nil {
		log.Fatal(err)
	}

	return
}

func (ctrl *DockerService) StartContainers(containerId string) {
	err := DockerClient.ContainerStart(context.Background(), containerId, types.ContainerStartOptions{})
	if err != nil {
		return
	}
}
