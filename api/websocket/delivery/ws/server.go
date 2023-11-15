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
	clients                   map[*WebsocketClient]bool
	topicHandlers             map[string]func() (interface{}, error)
	topicListeners            map[string][]*WebsocketClient
	topicsMessageForNewClient map[string]handler.ServerMessage
	topicHasClient            map[string]bool
}

type WebsocketClient struct {
	Connect *websocket.Conn
	Mutex   sync.Mutex
}

func (instance *WebsocketClient) SendJson(message handler.ServerMessage) error {
	instance.Mutex.Lock()
	defer instance.Mutex.Unlock()
	return instance.Connect.WriteJSON(message)
}

func NewWebsocket() *Ws {
	return &Ws{
		clients:                   make(map[*WebsocketClient]bool),
		topicHandlers:             make(map[string]func() (interface{}, error)),
		topicListeners:            make(map[string][]*WebsocketClient),
		topicsMessageForNewClient: make(map[string]handler.ServerMessage),
		topicHasClient:            make(map[string]bool),
	}
}

func (instance *Ws) SendMessage(topic string, message handler.ServerMessage) {

	for _, v := range instance.topicListeners[topic] {
		err := v.SendJson(message)
		if err != nil {
			fmt.Println("Can't send data to client ", v.Connect.RemoteAddr())
			continue
		} else {
		}
	}
}

func (instance *Ws) Handler() gin.HandlerFunc {
	return func(context *gin.Context) {

		connect, err := upgrader.Upgrade(context.Writer, context.Request, nil)

		wsClient := &WebsocketClient{
			Connect: connect,
		}

		defer connect.Close()
		defer delete(instance.clients, wsClient)
		defer instance.RemoveClientFromAllTopic(wsClient)
		defer fmt.Println("Client disconnected:", wsClient.Connect.RemoteAddr(), "Count clients:", len(instance.clients))

		instance.clients[wsClient] = true

		if err != nil {
			return
		}

		for {
			message := handler.Message{}
			err := connect.ReadJSON(&message)

			if err != nil {
				break
			}

			if message.EventType == "subscribe" {
				instance.AddClientToTopic(message.Subscription, wsClient)
			}
			if message.EventType == "unsubscribe" {
				instance.RemoveClientFromTopic(message.Subscription, wsClient)
			}
		}
	}
}

func (instance *Ws) AddClientToTopic(topic string, client *WebsocketClient) bool {
	for _, v := range instance.topicListeners[topic] {
		if v == client {
			return false
		}
	}

	instance.topicListeners[topic] = append(instance.topicListeners[topic], client)
	instance.topicHasClient[topic] = true

	client.SendJson(instance.CallHandler(topic))

	return true
}

func (instance *Ws) RemoveClientFromAllTopic(client *WebsocketClient) {
	for i, _ := range instance.topicListeners {
		instance.RemoveClientFromTopic(i, client)
	}
}

func (instance *Ws) RemoveClientFromTopic(topic string, client *WebsocketClient) bool {
	for i, clientItem := range instance.topicListeners[topic] {
		if clientItem == client {
			instance.topicListeners[topic] = append(instance.topicListeners[topic][:i], instance.topicListeners[topic][i+1:]...)

			if len(instance.topicListeners[topic]) <= 0 {
				instance.topicHasClient[topic] = false
			}
			return true
		}
	}
	return false
}

func (instance *Ws) HasTopicClients(topic string) bool {
	return instance.topicHasClient[topic]
}

func (instance *Ws) CallHandler(topic string) handler.ServerMessage {
	handlerMessage, err := instance.topicHandlers[topic]()
	serverMessage := handler.ServerMessage{Topic: topic, Data: handlerMessage}

	if err != nil {
		serverMessage.Error = gin.H{"serverMessage": err.Error()}
	}

	return serverMessage
}

func (instance *Ws) CreateTopic(topic string, fn func() (interface{}, error), params handler.HandlerParams) {
	_, ok := instance.topicListeners[topic]

	if !ok {
		instance.topicListeners[topic] = make([]*WebsocketClient, 0)
		instance.topicHandlers[topic] = fn
	}

	if params.SendOnce {
		return
	}
	go func() {
		for {
			instance.SendMessage(topic, instance.CallHandler(topic))

			time.Sleep(params.Sleep)
		}
	}()
}
