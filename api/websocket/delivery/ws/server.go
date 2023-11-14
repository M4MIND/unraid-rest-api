package ws

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"sync"
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
	clients map[*WebsocketClient]bool
	topics  map[string][]*WebsocketClient
}

type Message struct {
	EventType    string `json:"eventType"`
	Subscription string `json:"subscription"`
}

type WebsocketClient struct {
	Connect *websocket.Conn
	Mutex   sync.Mutex
}

func (class *WebsocketClient) SendJson(message handler.ServerMessage) error {
	class.Mutex.Lock()
	defer class.Mutex.Unlock()
	return class.Connect.WriteJSON(message)
}

func NewWebsocket() *Ws {
	return &Ws{
		clients: make(map[*WebsocketClient]bool),
		topics:  make(map[string][]*WebsocketClient),
	}
}

func (s *Ws) SendMessage(topic string, message handler.ServerMessage) {
	fmt.Println(topic, "Try send message")

	for _, v := range s.topics[topic] {
		err := v.SendJson(message)
		if err != nil {
			return
		} else {
			fmt.Println("Send Message", v.Connect.RemoteAddr())
		}
	}
}

func (s *Ws) Handler() gin.HandlerFunc {
	return func(context *gin.Context) {

		connect, err := upgrader.Upgrade(context.Writer, context.Request, nil)

		wsClient := &WebsocketClient{
			Connect: connect,
		}

		defer connect.Close()

		defer delete(s.clients, wsClient)

		s.clients[wsClient] = true

		if err != nil {
			return
		}

		for {
			message := Message{}
			err := connect.ReadJSON(&message)

			if err != nil {
				s.RemoveClientFromAllTopic(wsClient)
				fmt.Println("Client disconnected:", wsClient.Connect.RemoteAddr(), "Count clients:", len(s.clients)-1)
				break
			}

			if message.EventType == "subscribe" {
				s.AddClientToTopic(message.Subscription, wsClient)
			}
			if message.EventType == "unsubscribe" {
				s.RemoveClientFromTopic(message.Subscription, wsClient)
			}
		}

	}
}

func (s *Ws) AddClientToTopic(topic string, client *WebsocketClient) bool {
	for _, v := range s.topics[topic] {
		if v == client {
			return false
		}
	}

	s.topics[topic] = append(s.topics[topic], client)

	return true
}

func (s *Ws) RemoveClientFromAllTopic(client *WebsocketClient) {
	for i, _ := range s.topics {
		s.RemoveClientFromTopic(i, client)
	}
}

func (s *Ws) RemoveClientFromTopic(topic string, client *WebsocketClient) bool {
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

func (s *Ws) CreateTopic(topic string, fn func() handler.ServerMessage, sleep time.Duration) {
	_, ok := s.topics[topic]

	if !ok {
		s.topics[topic] = make([]*WebsocketClient, 0)
	}

	go func() {
		for {
			message := fn()
			message.Topic = topic
			s.SendMessage(topic, message)

			time.Sleep(sleep)
		}
	}()
}
