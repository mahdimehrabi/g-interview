package message

import (
	"errors"
	"github.com/rs/zerolog/log"

	"github.com/mahdimehrabi/graph-interview/receiver/internal/entity"
	"github.com/mahdimehrabi/graph-interview/receiver/internal/repository/message"
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
		log.Print("error saving message:", err.Error())
		return ErrMessageInternal
	}
	return nil
}
