package disks

import (
	"unraid-rest-api/pkg/monitoring"
	"unraid-rest-api/pkg/service"
)

type Monitor struct {
	services service.Container
}

func (d Monitor) Work() {

}

func New(s service.Container) monitoring.Monitoring {
	return &Monitor{services: s}
}
