package message

import (
	"errors"
	"fmt"

	"github.com/mahdimehrabi/graph-interview/broker/internal/entity"
	"github.com/mahdimehrabi/graph-interview/broker/internal/repository/message"
)

var ErrMessageInternal = errors.New("error")

type Message struct {
	messageRepo message.Message
}

func NewMessage(messageRepo message.Message) *Message {
	return &Message{
		messageRepo: messageRepo,
	}
}

func (m Message) Save(msg *entity.Message) error {
	if err := m.messageRepo.Save(msg); err != nil {
		fmt.Printf("error happend in sending message to destination: %s", err.Error())
		return ErrMessageInternal
	}
	return nil
}
