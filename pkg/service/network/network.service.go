package network

import (
	"time"
	"unraid-rest-api/pkg/service/network/types"

	"github.com/rafacas/sysstats"
)

type Service struct {
	history []types.NetworkAvg
	count   int
	max     int
}

func NewService() *Service {
	instance := &Service{max: 60}

	go instance._go()

	return instance
}

func (ctrl *Service) _go() {
	for {
		avg, _ := sysstats.GetNetStatsInterval(1)
		wrapper := types.NetworkAvg{Time: time.Now().UTC(), Avg: avg}
		if ctrl.count < ctrl.max {
			ctrl.history = append(ctrl.history, wrapper)
			ctrl.count++
		} else {
			ctrl.history = append(ctrl.history[1:], wrapper)
		}
	}
}

func (ctrl *Service) GetHistory() []types.NetworkAvg {
	return ctrl.history
}

func (ctrl *Service) GetLastHistory() types.NetworkAvg {
	return ctrl.history[ctrl.count-1]
}
