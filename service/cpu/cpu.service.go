package cpu

import (
	"time"
	"unraid-rest-api/service"
	"unraid-rest-api/service/cpu/types"

	"github.com/rafacas/sysstats"
	"github.com/shirou/gopsutil/cpu"
)

type Service struct {
	Container   service.Container
	countRepeat int
	maxHistory  int
	avgHistory  []types.CpuAvg
}

func NewService() *Service {
	instance := &Service{countRepeat: 0, maxHistory: 60}

	go instance._go()

	return instance
}

func (c *Service) _go() {
	for {
		avg, _ := sysstats.GetCpuStatsInterval(1)

		cpuAvgInstance := types.CpuAvg{
			Avg:  avg,
			Time: time.Now().UTC(),
		}

		if c.countRepeat < c.maxHistory {
			c.avgHistory = append(c.avgHistory, cpuAvgInstance)
			c.countRepeat++
		} else {
			c.avgHistory = append(c.avgHistory[1:], cpuAvgInstance)
		}
	}
}

func (c *Service) GetAvgHistory() []types.CpuAvg {
	return c.avgHistory
}

func (c *Service) GetAvgHistoryLast() types.CpuAvg {
	var lastIndex = len(c.avgHistory) - 1
	if lastIndex < 0 {
		avg, _ := sysstats.GetCpuStatsInterval(1)

		return types.CpuAvg{
			Avg:  avg,
			Time: time.Now().UTC(),
		}
	}
	return c.avgHistory[c.countRepeat-1]
}

func (c *Service) GetCpuInfo() ([]cpu.InfoStat, error) {
	return cpu.Info()
}
