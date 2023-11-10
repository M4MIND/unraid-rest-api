package cpu

import (
	"time"

	"github.com/rafacas/sysstats"
	"github.com/shirou/gopsutil/cpu"
)

type CpuAvg struct {
	Avg  sysstats.CpusAvgStats `json:"average"`
	Time time.Time             `json:"time"`
}

type CpuSysstats struct {
	countRepeat int
	maxHistory  int
	avgHistory  []CpuAvg
}

func NewCpuSysstats() *CpuSysstats {
	instance := &CpuSysstats{countRepeat: 0, maxHistory: 60}

	go instance._go()

	return instance
}

func (c *CpuSysstats) _go() {
	for {
		avg, _ := sysstats.GetCpuStatsInterval(1)

		cpuAvgInstance := CpuAvg{
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

func (c *CpuSysstats) GetAvgHistory() []CpuAvg {
	return c.avgHistory
}

func (c *CpuSysstats) GetAvgHistoryLast() CpuAvg {
	var lastIndex = len(c.avgHistory) - 1
	if lastIndex < 0 {
		avg, _ := sysstats.GetCpuStatsInterval(1)

		return CpuAvg{
			Avg:  avg,
			Time: time.Now().UTC(),
		}
	}
	return c.avgHistory[c.countRepeat-1]
}

func (c *CpuSysstats) GetCpuInfo() ([]cpu.InfoStat, error) {
	return cpu.Info()
}
