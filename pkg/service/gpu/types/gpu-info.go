package types

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
