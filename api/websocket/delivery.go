package websocket

type Handlers interface {
	CpuState() ServerMessage
	PingPong() ServerMessage
}

type ServerMessage struct {
	Data  interface{} `json:"data"`
	Topic string      `json:"topic"`
}
