package memory

import (
	"time"

	"github.com/rafacas/sysstats"
)

type Service struct {
	history     []Memory
	countRepeat int
	maxHistory  int
}

type Memory struct {
	Stats sysstats.MemStats `json:"stats"`
	Time  time.Time         `json:"time"`
}

func NewService() *Service {
	instance := &Service{maxHistory: 60}

	go instance._go()

	return instance
}

func (c *Service) _go() {
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

func (c *Service) GetHistory() []Memory {
	return c.history
}

func (c *Service) GetHistoryLast() Memory {
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
