package websocket

type ServerMessage struct {
	Data  interface{} `json:"data"`
	Topic string      `json:"topic"`
	Error interface{} `json:"error"`
}

type ServerMessageErrorData struct {
	Message string `json:"message"`
}

type Message struct {
	EventType    string `json:"eventType"`
	Subscription string `json:"subscription"`
}
