package services

import (
	"time"

	"github.com/rafacas/sysstats"
)

type MemoryService struct {
	lastTime int64
	count    int
	history  [60]sysstats.MemStats
	last     sysstats.MemStats
}

func (ctrl *MemoryService) Go() {
	ctrl.lastTime = time.Now().Unix()

	for true {
		now := time.Now()
		if now.Sub(time.Unix(ctrl.lastTime, 0)).Seconds() < 1 {
			continue
		}

		current := ctrl.GetMemoryInfo()
		if ctrl.count < len(ctrl.history) {
			ctrl.history[ctrl.count] = current
			ctrl.count++

		} else {
			temp := ctrl.history[1:]
			ctrl.history = [60]sysstats.MemStats(append(temp, current))
		}

		ctrl.last = current
		ctrl.lastTime = now.Unix()
	}
}

func (ctrl *MemoryService) GetMemoryInfo() sysstats.MemStats {
	stat, _ := sysstats.GetMemStats()
	return stat
}

func (ctrl *MemoryService) GetMemoryHistory() [60]sysstats.MemStats {
	return ctrl.history
}
