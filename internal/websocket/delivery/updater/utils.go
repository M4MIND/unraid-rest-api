package updater

import "github.com/gorilla/websocket"

func FindConnection(clientsToSubscripctions map[*websocket.Conn]map[string]bool, subscriptionType string) *websocket.Conn {
	if len(clientsToSubscripctions) < 1 {
		return nil
	}
	for clientConnection, subscription := range clientsToSubscripctions {
		_, ok := subscription[subscriptionType]
		if ok {
			return clientConnection
		}
	}
	return nil
}
