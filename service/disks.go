package service

import (
	"time"
	"unraid-rest-api/service/utils"
)

type DisksSysstats struct {
	history []Stat
	count   int
	max     int
	parser  utils.DisksUtils
}

type Stat struct {
	Time time.Time
	Avg  []utils.DiskAvgStats
}

type DiskSysstatsRaw struct {
	Major        int    `json:"major"`        // Major number for the disk
	Minor        int    `json:"minor"`        // Minor number for the disk
	Name         string `json:"name"`         // Disk name
	ReadIOs      uint64 `json:"readios"`      // # of reads completed since boot
	ReadMerges   uint64 `json:"readmerges"`   // # of reads merged since boot
	ReadSectors  uint64 `json:"readsectors"`  // # of sectors read since boot
	ReadTicks    uint64 `json:"readticks"`    // # of milliseconds spent reading since boot
	WriteIOs     uint64 `json:"writeios"`     // # of writes completed since boot
	WriteMerges  uint64 `json:"writemerges"`  // # of writes merged since boot
	WriteSectors uint64 `json:"writesectors"` // # of sectors written since boot
	WriteTicks   uint64 `json:"writeticks"`   // # of milliseconds spent writing since boot
	InFlight     uint64 `json:"inflight"`     // # of I/Os currently in progress
	IOTicks      uint64 `json:"ioticks"`      // # of milliseconds spent doing I/Os since boot
	TimeInQueue  uint64 `json:"timeinqueue"`  // Weighted # of milliseconds spent doing I/Os since boot
	SampleTime   int64  `json:"sampletime"`   // Time when the sample was taken
}

func NewDisksSysstats() *DisksSysstats {
	instance := &DisksSysstats{max: 60}
	instance.parser = utils.NewDiskUtils()

	go instance._go()

	return instance
}

func (s *DisksSysstats) _go() {
	for {

		avg := s.parser.GetAvgStatsInterval(1)
		wrapper := Stat{Avg: avg, Time: time.Now().UTC()}

		if s.count < s.max {
			s.history = append(s.history, wrapper)
			s.count++
		} else {
			s.history = append(s.history[1:], wrapper)
		}
	}
}

func (s *DisksSysstats) GetHistory() []Stat {
	return s.history
}

func (s *DisksSysstats) GetHistoryLast() Stat {
	return s.history[s.count-1]
}
