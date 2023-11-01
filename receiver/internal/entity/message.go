package entity

import (
	"github.com/google/uuid"
	"time"
)

type Message struct {
	Message    string
	ReceivedAt int64
	ID         uuid.UUID // to get response base on this id and check whether it successfully inserted or not
}

func NewMessage(msg string) *Message {
	return &Message{
		ReceivedAt: time.Now().Unix(),
		Message:    msg,
	}
}
