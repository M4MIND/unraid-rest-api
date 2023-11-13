package ws

import (
	"time"
	"unraid-rest-api/api/websocket"
	"unraid-rest-api/pkg/service"
)

type Handler struct {
	service service.Container
}

func (h Handler) CpuState() websocket.ServerMessage {
	time.Sleep(1 * time.Second)
	return websocket.ServerMessage{
		Data: []byte("asd"),
	}
}

func NewHandler(s service.Container) websocket.Handlers {
	return &Handler{service: s}
}
