package entity

import "time"

type Message struct {
	Message    string
	ReceivedAt int
	// extra entity props
}

func NewMessage(msg string) *Message {
	return &Message{
		ReceivedAt: int(time.Now().Unix()),
		Message:    msg,
	}
}
