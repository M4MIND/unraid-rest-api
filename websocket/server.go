package websocket

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

type Ws struct {
	Clients map[uuid.UUID]*Client
	Hubs    map[string]*Hub
}

func NewServer() *Ws {
	instance := &Ws{
		Clients: make(map[uuid.UUID]*Client),
		Hubs:    make(map[string]*Hub),
	}

	instance.CreateHub("DisksArrayInfo")

	return instance
}

func (w *Ws) HandleSubscribeMessage(message *SubscribeMessage, client *Client) bool {
	return w.SubscribeHub(message.Hub, client)
}

func (w *Ws) HandleUnsubscribeMessage(message *SubscribeMessage, client *Client) bool {
	return w.UnsubscribeHub(message.Hub, client)
}

func (w *Ws) BreakConnection(client *Client) {
	for _, v := range client.Hubs {
		v.Unsubscribe(client)
	}

	delete(w.Clients, client.UUID)
}

func (w *Ws) Handler(c *gin.Context) {
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)

	if err != nil {
		fmt.Println(err)

		return
	}

	client := &Client{
		UUID:       uuid.New(),
		Connection: ws, Hubs: make([]Hub, 0),
		CallbackSubscribe:       w.HandleSubscribeMessage,
		CallbackUnsubscribe:     w.HandleUnsubscribeMessage,
		CallbackBreakConnection: w.BreakConnection,
	}

	w.Clients[client.UUID] = client

	go client.Listen()
}

func (w *Ws) SubscribeHub(name string, client *Client) bool {
	_, ok := w.Hubs[name]

	if ok {
		w.Hubs[name].Subscribe(client)
	}

	return ok
}

func (w *Ws) UnsubscribeHub(name string, client *Client) bool {
	_, ok := w.Hubs[name]

	if ok {
		return w.Hubs[name].Unsubscribe(client)
	}

	return false
}

func (w *Ws) CreateHub(name string) *Hub {
	_, ok := w.Hubs[name]

	if ok {
		panic("Hub is created {" + name + "}")
	}

	w.Hubs[name] = &Hub{
		Clients: make(map[uuid.UUID]*Client),
	}

	return w.Hubs[name]
}
