package types

type Message struct {
	Type string `json:"type"`
}

type ConsumeMessage struct {
	Message
	CursorName string `json:"cursor_name"`
	EventType  string `json:"event_type"`
}

type ConsumeMessageResponse struct {
	EventID   string `json:"event_id"`
	EventData string `json:"event_data"`
}

type AckMessage struct {
	Message
	CursorName string `json:"cursor_name"`
	EventID    string `json:"event_id"`
}

type AckMessageResponse struct {
	Status int `json:"status"`
}
