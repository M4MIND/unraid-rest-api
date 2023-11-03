package service

import "github.com/rafacas/sysstats"

type NetworkSysstats struct {
	history []sysstats.NetAvgStats
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
		if ctrl.count < ctrl.max {
			ctrl.history = append(ctrl.history, avg)
			ctrl.count++
		} else {
			ctrl.history = append(ctrl.history[1:], avg)
		}
	}
}

func (ctrl *NetworkSysstats) GetHistory() []sysstats.NetAvgStats {
	return ctrl.history
}

func (ctrl *NetworkSysstats) Get() sysstats.NetRawStats {
	stat, _ := sysstats.GetNetRawStats()

	return stat
}
