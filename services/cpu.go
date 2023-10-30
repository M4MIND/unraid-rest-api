package services

import (
	"github.com/rafacas/sysstats"
	"log"
)

type CpuService struct {
	lastTime         int64
	countRepeat      int
	cpusAvgStats     [60]sysstats.CpusAvgStats
	lastCpusRawState sysstats.CpusRawStats
}

func (ctrl *CpuService) Go() {
	for true {
		avg, _ := sysstats.GetCpuStatsInterval(1)

		if ctrl.countRepeat != len(ctrl.cpusAvgStats)-1 {
			ctrl.countRepeat++
			ctrl.cpusAvgStats[ctrl.countRepeat] = avg
		} else {
			temp := ctrl.cpusAvgStats[1:]
			ctrl.cpusAvgStats = [60]sysstats.CpusAvgStats(append(temp, avg))
		}
	}
}

func getCpuRawStats() sysstats.CpusRawStats {
	stats, err := sysstats.GetCpuRawStats()

	if err != nil {
		log.Fatal(err)
	}

	return stats

}

func (ctrl *CpuService) GetAvgHistory() [60]sysstats.CpusAvgStats {
	return ctrl.cpusAvgStats
}

func (ctrl *CpuService) GetLastCpuAvg() sysstats.CpusAvgStats {
	return ctrl.cpusAvgStats[ctrl.countRepeat]
}
