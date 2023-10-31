package service

import "github.com/rafacas/sysstats"

type NetworkSysstats struct {
	history [60]sysstats.NetAvgStats
}

func NewNetworkSysstats() *NetworkSysstats {
	instance := &NetworkSysstats{}

	go instance._go()

	return instance
}

func (ctrl *NetworkSysstats) _go() {
	count := 0
	for {
		avg, _ := sysstats.GetNetStatsInterval(1)

		if count < len(ctrl.history) {
			ctrl.history[count] = avg
			count++
		} else {
			temp := ctrl.history[1:]
			ctrl.history = [60]sysstats.NetAvgStats(append(temp, avg))
		}
	}
}

func (ctrl *NetworkSysstats) GetNetworkStats() {

}

func (ctrl *NetworkSysstats) GetHistory() [60]sysstats.NetAvgStats {
	return ctrl.history
}

func (ctrl *NetworkSysstats) Get() sysstats.NetRawStats {
	stat, _ := sysstats.GetNetRawStats()

	return stat
}
