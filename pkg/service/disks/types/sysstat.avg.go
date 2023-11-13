package types

type SysstatAvg struct {
	Major       int     `json:"major"`       // Major number for the disk
	Minor       int     `json:"minor"`       // Minor number for the disk
	Name        string  `json:"name"`        // Disk name
	ReadIOs     float64 `json:"readios"`     // # of reads completed per second
	ReadMerges  float64 `json:"readmerges"`  // # of reads merged per second
	ReadBytes   float64 `json:"readbytes"`   // # of bytes read per second
	WriteIOs    float64 `json:"writeios"`    // # of writes completed per second
	WriteMerges float64 `json:"writemerges"` // # of writes merged per second
	WriteBytes  float64 `json:"writebytes"`  // # of bytes written per second
	InFlight    uint64  `json:"inflight"`    // # of I/Os currently in progress
	IOTicks     uint64  `json:"ioticks"`     // # of milliseconds spent doing I/Os
	TimeInQueue uint64  `json:"timeinqueue"` // Weighted # of milliseconds spent doing I/Os
}
