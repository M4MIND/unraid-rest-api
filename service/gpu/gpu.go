package gpu

type GpuInfoInstance interface {
	GetInfo() (stats GpuInfo, status bool)
	GetStatus() bool
}

type GpuInfo struct {
	Model string
}
