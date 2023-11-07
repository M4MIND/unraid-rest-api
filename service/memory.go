package service

import (
	"time"

	"github.com/rafacas/sysstats"
)

type MemorySysstats struct {
	history     []Memory
	countRepeat int
	maxHistory  int
}

type Memory struct {
	Stats sysstats.MemStats `json:"stats"`
	Time  time.Time         `json:"time"`
}

func NewMemorySysstats() *MemorySysstats {
	instance := &MemorySysstats{maxHistory: 60}

	go instance._go()

	return instance
}

func (c *MemorySysstats) _go() {
	for {
		time.Sleep(1 * time.Second)

		stat, _ := sysstats.GetMemStats()

		instance := Memory{Stats: stat, Time: time.Now().UTC()}

		if c.countRepeat < c.maxHistory {
			c.history = append(c.history, instance)
			c.countRepeat++
		} else {
			c.history = append(c.history[1:], instance)
		}
	}
}

func (c *MemorySysstats) GetHistory() []Memory {
	return c.history
}

func (c *MemorySysstats) GetHistoryLast() Memory {
	var lastIndex = len(c.history) - 1
	if lastIndex < 0 {
		stat, _ := sysstats.GetMemStats()

		return Memory{
			Stats: stat,
			Time:  time.Now().UTC(),
		}
	}
	return c.history[lastIndex]
}
