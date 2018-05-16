type RequestMessage struct {
	SenderID  ClientID
	HandlerID MessageHandlerID
	Body      []byte
}

type ResponseMessage struct {
	SenderID  ClientID
	EventType EventType
	HandlerID MessageHandlerID
	CastType  CastType
	CastFor   []ClientID
	Body      []byte
}
