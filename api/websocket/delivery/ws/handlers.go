package ws

import (
	"errors"
	"github.com/gin-gonic/gin"
	"unraid-rest-api/api/websocket"
	"unraid-rest-api/pkg/service"
)

type Handler struct {
	service service.Container
}

func (h Handler) PingPong() (interface{}, error) {
	return gin.H{"Ping": "Pong"}, nil
}

func (h Handler) CpuState() (interface{}, error) {

	return nil, errors.New("ERRROR")
	//return h.service.CpuService.GetAvgHistoryLast(), nil
}

func NewHandler(s service.Container) websocket.Handlers {
	return &Handler{service: s}
}
