package types

type SysstatRaw struct {
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
