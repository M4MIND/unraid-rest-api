package types

type GpuInfoInstance interface {
	GetInfo() (stats GpuInfo, status bool)
	GetStatus() bool
}
