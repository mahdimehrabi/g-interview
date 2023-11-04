package message

import (
	"errors"

	"github.com/mahdimehrabi/graph-interview/broker/internal/entity"
	"github.com/mahdimehrabi/graph-interview/broker/internal/repository/message"
	"github.com/rs/zerolog/log"
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
		log.Printf("error happened in sending message to destination: %s", err.Error())
		return ErrMessageInternal
	}
	return nil
}
