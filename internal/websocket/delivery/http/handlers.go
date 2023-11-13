package http

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"
	"unraid-rest-api/internal/websocket"
	"unraid-rest-api/internal/websocket/delivery/message"
	"unraid-rest-api/internal/websocket/delivery/updater"
	"unraid-rest-api/service"
	"unraid-rest-api/service/cpu/types"

	"github.com/gin-gonic/gin"

	"net/http"

	gorillaWS "github.com/gorilla/websocket"
)

type WebsocketHandler struct {
	services service.Container
}

var clientsToSubscripctions = make(map[*gorillaWS.Conn]map[string]bool)
var messages = []message.ServerMessage{}

func NewHandler(services service.Container) websocket.Handlers {
	handler := &WebsocketHandler{services: services}

	go beginPingPong(&services)
	go beginUpdateCpu(&services)

	return handler
}

var upgrader = gorillaWS.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// Allow all connections
		return true
	},
}

func (s *WebsocketHandler) UpgradeWebsocket() gin.HandlerFunc {
	return func(context *gin.Context) {
		conn, err := upgrader.Upgrade(context.Writer, context.Request, nil)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		go s.handleWebSocketConnection(conn)
		go s.sendWebsocketMessages()
	}

}

type WSClientMessage struct {
	EventType    string `json:"eventType"`
	Subscription string `json:"subscription"`
}

func (s *WebsocketHandler) handleWebSocketConnection(conn *gorillaWS.Conn) {
	for {
		// Read message from the client
		var message WSClientMessage

		err := conn.ReadJSON(&message)
		if err != nil {
			fmt.Printf("Something went wrong: %s", err.Error())
			conn.Close()
			delete(clientsToSubscripctions, conn)
			break
		}

		fmt.Println("test1")
		fmt.Printf("%+v\n", message)

		clientSubscripctions := clientsToSubscripctions[conn]

		if clientSubscripctions == nil {
			clientSubscripctions = make(map[string]bool)
		}

		clientsToSubscripctions[conn] = clientSubscripctions
		clientSubscripctions[message.Subscription] = true

	}
}

func (s *WebsocketHandler) sendWebsocketMessages() {
	for {
		for i, message := range messages {
			message = messages[i]
			data := message.Data
			connection := message.Connection
			mutex := &messages[i].Mutex

			mutex.Lock()
			defer mutex.Unlock()
			err := connection.WriteMessage(1, data)

			if err != nil {
				connection.Close()
				delete(clientsToSubscripctions, connection)
			}
			messages = append(messages[:0], messages[1:]...)

		}

	}
}

type CpuDataPoint struct {
	SubscriptionType string       `json:"subscriptionType"`
	Time             time.Time    `json:"time"`
	Value            types.CpuAvg `json:"value"`
}

const cpuSubscriptionType = "cpu-data"
const tickRateSec = 1

func beginUpdateCpu(s *service.Container) {
	for {
		clientConnection := updater.FindConnection(clientsToSubscripctions, cpuSubscriptionType)
		if clientConnection != nil {

			dataPoint := CpuDataPoint{
				SubscriptionType: cpuSubscriptionType,
				Time:             time.Now(),
				Value:            s.CpuService.GetAvgHistoryLast(),
			}

			data, err := json.Marshal(dataPoint)
			if err != nil {
				fmt.Println(err)
				return
			}
			messages = append(messages, message.ServerMessage{
				Connection: clientConnection,
				Data:       data,
				Mutex:      sync.Mutex{},
			})

		}

		time.Sleep(tickRateSec * time.Second)
	}
}

type PongDataPoint struct {
	SubscriptionType string    `json:"subscriptionType"`
	Time             time.Time `json:"time"`
	Value            string    `json:"value"`
}

const pongSubscriptionType = "ping-pong"
const pongTickRateSec = 1

func beginPingPong(s *service.Container) {
	for {
		clientConnection := updater.FindConnection(clientsToSubscripctions, pongSubscriptionType)
		if clientConnection != nil {

			dataPoint := PongDataPoint{
				SubscriptionType: pongSubscriptionType,
				Time:             time.Now(),
				Value:            "pong",
			}

			data, err := json.Marshal(dataPoint)
			if err != nil {
				fmt.Println(err)
				return
			}
			messages = append(messages, message.ServerMessage{
				Connection: clientConnection,
				Data:       data,
				Mutex:      sync.Mutex{},
			})
		}

		time.Sleep(pongTickRateSec * time.Second)
	}
}
