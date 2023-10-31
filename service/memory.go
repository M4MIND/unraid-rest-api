package service

import (
	"github.com/rafacas/sysstats"
	"time"
)

type MemorySysstats struct {
	history     [60]sysstats.MemStats
	countRepeat int
}

func NewMemorySysstats() *MemorySysstats {
	instance := &MemorySysstats{}

	go instance._go()

	return instance
}

func (c *MemorySysstats) _go() {
	for {
		time.Sleep(1 * time.Second)

		stat, _ := sysstats.GetMemStats()

		if c.countRepeat != len(c.history)-1 {
			c.history[c.countRepeat] = stat
			c.countRepeat++
		} else {
			temp := c.history[1:]
			c.history = [60]sysstats.MemStats(append(temp, stat))
		}
	}
}

func (c *MemorySysstats) GetHistory() [60]sysstats.MemStats {
	return c.history
}

func (c *MemorySysstats) GetHistoryLast() sysstats.MemStats {
	return c.history[c.countRepeat]
}
