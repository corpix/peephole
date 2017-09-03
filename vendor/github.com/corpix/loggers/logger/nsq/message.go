package nsq

// Message is a logger message data.
type Message struct {
	Level   string
	Payload interface{}
}

// NewMessage creates new Message with logging level as string and
// payload in original type.
func NewMessage(lvl Level, payload interface{}) *Message {
	return &Message{lvl.String(), payload}
}
