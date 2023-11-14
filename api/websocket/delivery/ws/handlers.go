package ws

import (
	"unraid-rest-api/api/websocket"
	"unraid-rest-api/pkg/service"
)

type Handler struct {
	service service.Container
}

func (h Handler) PingPong() websocket.ServerMessage {
	return websocket.ServerMessage{
		Data: []byte("Ping"),
	}
}

func (h Handler) CpuState() websocket.ServerMessage {

	return websocket.ServerMessage{
		Data: h.service.CpuService.GetAvgHistoryLast(),
	}
}

func NewHandler(s service.Container) websocket.Handlers {
	return &Handler{service: s}
}
