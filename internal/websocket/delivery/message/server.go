package message

import (
	"sync"

	"github.com/gorilla/websocket"
)

type ServerMessage struct {
	Connection *websocket.Conn
	Data       []byte
	Mutex      sync.Mutex
}
