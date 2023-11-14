package websocket

type Handlers interface {
	CpuState() (interface{}, error)
	PingPong() (interface{}, error)
	ArrayInfo() (interface{}, error)
	MemoryInfoTick() (interface{}, error)
}
