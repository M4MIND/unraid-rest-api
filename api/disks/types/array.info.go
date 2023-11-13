package types

type ArrayInfo struct {
	Size    int64
	Used    int64
	Free    int64
	Devices []ArrayInfoDevice
}

type ArrayInfoDevice struct {
	DiskId          string
	DiskSizeBytes   int64
	DiskState       int
	RdevName        string
	DiskUsedPercent float32
	DiskUsedBytes   int64
	Temperature     int
	IsHdd           bool
	Mount           string
}
