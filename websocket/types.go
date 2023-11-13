package websocket

import (
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type StatusMessage struct {
	Message     string `json:"message"`
	Status      int    `json:"status"`
	Description string `json:"description"`
}

type Client struct {
	UUID                    uuid.UUID
	Connection              *websocket.Conn
	Hubs                    []Hub
	CallbackSubscribe       func(message *SubscribeMessage, client *Client) bool
	CallbackUnsubscribe     func(message *SubscribeMessage, client *Client) bool
	CallbackBreakConnection func(client *Client)
}
type SubscribeMessage struct {
	Hub       string `json:"hub"`
	Subscribe bool   `json:"subscribe"`
}

func (client *Client) SendJson(v interface{}) {
	err := client.Connection.WriteJSON(v)
	if err != nil {
		return
	}
}

func (client *Client) Listen() {
	for {
		message := &SubscribeMessage{}
		//Read Message from client
		err := client.Connection.ReadJSON(message)

		if err != nil {
			client.CallbackBreakConnection(client)
			break
		}

		if message.Subscribe {
			if !client.CallbackSubscribe(message, client) {
				client.SendJson(StatusMessage{
					Status:      500,
					Message:     "Failed Subscribe",
					Description: message.Hub,
				})
			} else {
				client.SendJson(StatusMessage{
					Status:      200,
					Message:     "Success subscribe",
					Description: message.Hub,
				})
			}
		} else {
			if !client.CallbackUnsubscribe(message, client) {
				client.SendJson(StatusMessage{
					Status:      500,
					Message:     "Failed unsubscribe",
					Description: message.Hub,
				})
			} else {
				client.SendJson(StatusMessage{
					Status:      200,
					Message:     "Success unsubscribe",
					Description: message.Hub,
				})
			}
		}
	}

	defer client.Connection.Close()
}

type Hub struct {
	Clients map[uuid.UUID]*Client
}

func (h *Hub) Subscribe(client *Client) bool {
	_, ok := h.Clients[client.UUID]

	if !ok {
		h.Clients[client.UUID] = client
	}

	return ok
}

func (h *Hub) Unsubscribe(client *Client) bool {
	_, ok := h.Clients[client.UUID]

	if ok {
		delete(h.Clients, client.UUID)
	}

	return ok
}

func (h *Hub) BroadcastJson(message interface{}) {
	for _, v := range h.Clients {
		v.SendJson(message)
	}
}
