package message

import (
	"errors"

	"github.com/mahdimehrabi/graph-interview/destination/internal/entity"
	"github.com/mahdimehrabi/graph-interview/destination/internal/repository/message"
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
	if err := m.messageRepo.SaveAnalyze(msg); err != nil {
		return ErrMessageInternal
	}
	return nil
}
