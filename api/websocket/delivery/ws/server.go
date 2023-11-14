package ws

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
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
	topics  map[string][]*websocket.Conn
}

type Message struct {
	EventType    string `json:"eventType"`
	Subscription string `json:"subscription"`
}

func NewWebsocket() *Ws {
	return &Ws{
		clients: make(map[*websocket.Conn]bool),
		topics:  make(map[string][]*websocket.Conn),
	}
}

func (s *Ws) SendMessage(topic string, message handler.ServerMessage) {
	fmt.Println(topic, "Try send message")

	for _, v := range s.topics[topic] {
		fmt.Println("Send Message", v.RemoteAddr())
		err := v.WriteJSON(message)
		if err != nil {
			log.Println(err)
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
			return
		}

		for {
			message := Message{}
			err := connect.ReadJSON(&message)

			if err != nil {
				s.RemoveClientFromAllTopic(connect)
				break
			}

			if message.EventType == "subscribe" {
				s.AddClientToTopic(message.Subscription, connect)
			}
			if message.EventType == "unsubscribe" {
				s.RemoveClientFromTopic(message.Subscription, connect)
			}
		}

	}
}

func (s *Ws) AddClientToTopic(topic string, client *websocket.Conn) bool {
	for _, v := range s.topics[topic] {
		if v == client {
			return false
		}
	}

	s.topics[topic] = append(s.topics[topic], client)

	return true
}

func (s *Ws) RemoveClientFromAllTopic(client *websocket.Conn) {
	for i, _ := range s.topics {
		s.RemoveClientFromTopic(i, client)
	}
}

func (s *Ws) RemoveClientFromTopic(topic string, client *websocket.Conn) bool {
	for i, v := range s.topics[topic] {
		if v == client {
			s.topics[topic] = append(s.topics[topic][:i], s.topics[topic][i+1:]...)
			return true
		}
	}
	return false
}

func (s *Ws) HasTopicClients(topic string) bool {
	return len(s.topics[topic]) > 0
}

func (s *Ws) CreateTopic(topic string, fn func() handler.ServerMessage, sleepSeconds int64) {
	_, ok := s.topics[topic]

	if !ok {
		s.topics[topic] = make([]*websocket.Conn, 0)
	}

	go func() {
		for {
			s.SendMessage(topic, fn())

			time.Sleep(time.Duration(sleepSeconds) * time.Second)
		}
	}()
}
