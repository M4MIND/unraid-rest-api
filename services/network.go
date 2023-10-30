package services

import "github.com/rafacas/sysstats"

type NetworkService struct {
	history [60]sysstats.NetAvgStats
}

func (ctrl *NetworkService) Go() {
	count := 0
	for true {
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

func (ctrl *NetworkService) GetNetworkStats() {

}

func (ctrl *NetworkService) GetHistory() [60]sysstats.NetAvgStats {
	return ctrl.history
}

func (ctrl *NetworkService) Get() sysstats.NetRawStats {
	stat, _ := sysstats.GetNetRawStats()

	return stat
}
