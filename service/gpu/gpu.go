package gpu

type GpuInfoInstance interface {
	GetInfo() (stats GpuInfo, status bool)
	GetStatus() bool
}

type GpuInfo struct {
	Model         string
	DriverVersion string
	Fan           string
	MemReserved   string
	MemUsed       string
	MemFree       string
	MemTotal      string
	TempMemory    string
	TempGpu       string
}
