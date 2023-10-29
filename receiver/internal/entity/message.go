package entity

import "time"

type Message struct {
	Message    string
	ReceivedAt int64
	// extra entity props
}

func NewMessage(msg string) *Message {
	return &Message{
		ReceivedAt: time.Now().Unix(),
		Message:    msg,
	}
}
