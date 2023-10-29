package services

import (
	"github.com/rafacas/sysstats"
	"log"
	"time"
)

var lastTime = time.Now()
var cpusAvgStats [60]sysstats.CpusAvgStats
var countRepeat = 0
var lastCpusRawState sysstats.CpusRawStats

type CpuService struct {
}

func (ctrl *CpuService) Go() {
	lastTime = time.Now()

	lastCpusRawState = getCpuRawStats()

	for true {
		t := time.Now()

		if t.Sub(lastTime).Seconds() < 1 {
			continue
		}

		lastTime = t

		current := getCpuRawStats()
		avg, _ := sysstats.GetCpuAvgStats(lastCpusRawState, current)

		if countRepeat != len(cpusAvgStats)-1 {
			countRepeat++
			cpusAvgStats[countRepeat] = avg
		} else {
			temp := cpusAvgStats[1:]
			cpusAvgStats = [60]sysstats.CpusAvgStats(append(temp, avg))
		}

		lastCpusRawState = current
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
	return cpusAvgStats
}

func (ctrl *CpuService) GetLastCpuAvg() sysstats.CpusAvgStats {
	return cpusAvgStats[countRepeat]
}
