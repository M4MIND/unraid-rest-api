package service

import (
	"github.com/rafacas/sysstats"
	"github.com/shirou/gopsutil/cpu"
)

type CpuSysstats struct {
	countRepeat int
	avgHistory  [60]sysstats.CpusAvgStats
}

func CpuNew() *CpuSysstats {
	instance := &CpuSysstats{}

	go instance._go()

	return instance
}

func (c *CpuSysstats) _go() {
	for {
		avg, _ := sysstats.GetCpuStatsInterval(1)

		if c.countRepeat != len(c.avgHistory)-1 {
			c.countRepeat++
			c.avgHistory[c.countRepeat] = avg
		} else {
			temp := c.avgHistory[1:]
			c.avgHistory = [60]sysstats.CpusAvgStats(append(temp, avg))
		}
	}
}

func (c *CpuSysstats) GetAvgHistory() [60]sysstats.CpusAvgStats {
	return c.avgHistory
}

func (c *CpuSysstats) GetAvgHistoryLast() sysstats.CpusAvgStats {
	return c.avgHistory[c.countRepeat]
}

func (c *CpuSysstats) GetCpuInfo() ([]cpu.InfoStat, error) {
	return cpu.Info()
}
