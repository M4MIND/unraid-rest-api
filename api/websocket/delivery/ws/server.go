package ws

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	handler "unraid-rest-api/api/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// Allow all connections
		return true
	},
}

type Ws struct {
	clients map[*websocket.Conn]bool
	topics  map[string]*websocket.Conn
}

type Message struct {
	EventType    string `json:"eventType"`
	Subscription string `json:"subscription"`
}

func NewWebsocket() *Ws {
	return &Ws{
		clients: make(map[*websocket.Conn]bool),
		topics:  make(map[string]*websocket.Conn),
	}
}

func (s *Ws) SendMessage(topic string, message handler.ServerMessage) {
	for _, v := range s.topics {
		err := v.WriteJSON(message)
		if err != nil {
			return
		}
	}
}

func (s *Ws) Handler() gin.HandlerFunc {
	return func(context *gin.Context) {

		connect, err := upgrader.Upgrade(context.Writer, context.Request, nil)

		defer connect.Close()

		defer delete(s.clients, connect)

		s.clients[connect] = true

		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		for {
			message := Message{}
			err := connect.ReadJSON(&message)

			if err != nil {
				break // Выходим из цикла, если клиент пытается закрыть соединение или связь прервана
			}

			if message.EventType == "subscribe" {
				s.topics[message.Subscription] = connect
			}
		}

	}
}

func (s *Ws) AddTopic(topic string, fn func() handler.ServerMessage) {
	go func() {
		for {
			s.SendMessage(topic, fn())
		}
	}()
}
