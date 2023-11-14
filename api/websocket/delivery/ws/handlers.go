package ws

import (
	"encoding/json"
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

	data, _ := json.Marshal(h.service.CpuService.GetAvgHistoryLast())

	return websocket.ServerMessage{
		Data: data,
	}
}

func NewHandler(s service.Container) websocket.Handlers {
	return &Handler{service: s}
}
