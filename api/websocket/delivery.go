package websocket

import (
	"github.com/gorilla/websocket"
	"sync"
)

type Handlers interface {
	CpuState() ServerMessage
	PingPong() ServerMessage
}

type ServerMessage struct {
	Connection *websocket.Conn
	Data       []byte
	Mutex      sync.Mutex
}
