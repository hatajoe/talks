package hey
...
type OnReceive interface {
	Receive(string)
}

func Connect(url, origin string, h OnReceive) {
	var err error
	ws, err = websocket.Dial(url, "", origin)
	if err != nil {
		panic(err)
	}
	go receive(h)
}

func receive(h OnReceive) {
	var msg broadcaster.ResponseMessage
	for {
		if err := websocket.JSON.Receive(ws, &msg); err != nil {
			panic(err)
		}
		h.Receive(string(msg.Body))
	}
}
