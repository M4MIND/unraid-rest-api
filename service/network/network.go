package network

import (
	"github.com/rafacas/sysstats"
	"time"
)

type NetworkAvg struct {
	Avg  sysstats.NetAvgStats
	Time time.Time
}

type NetworkSysstats struct {
	history []NetworkAvg
	count   int
	max     int
}

func NewNetworkSysstats() *NetworkSysstats {
	instance := &NetworkSysstats{max: 60}

	go instance._go()

	return instance
}

func (ctrl *NetworkSysstats) _go() {
	for {
		avg, _ := sysstats.GetNetStatsInterval(1)
		wrapper := NetworkAvg{Time: time.Now().UTC(), Avg: avg}
		if ctrl.count < ctrl.max {
			ctrl.history = append(ctrl.history, wrapper)
			ctrl.count++
		} else {
			ctrl.history = append(ctrl.history[1:], wrapper)
		}
	}
}

func (ctrl *NetworkSysstats) GetHistory() []NetworkAvg {
	return ctrl.history
}

func (ctrl *NetworkSysstats) GetLastHistory() NetworkAvg {
	return ctrl.history[ctrl.count-1]
}

func (ctrl *NetworkSysstats) Get() sysstats.NetRawStats {
	stat, _ := sysstats.GetNetRawStats()

	return stat
}
