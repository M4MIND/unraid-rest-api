package websocket

type Handlers interface {
	CpuState() (interface{}, error)
	PingPong() (interface{}, error)
}

type ServerMessage struct {
	Data  interface{} `json:"data"`
	Topic string      `json:"topic"`
	Error interface{} `json:"error"`
}

type ServerMessageErrorData struct {
	Message string `json:"message"`
}
