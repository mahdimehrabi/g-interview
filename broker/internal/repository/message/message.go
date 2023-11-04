package message

import "github.com/mahdimehrabi/graph-interview/broker/internal/entity"

type Message interface {
	Save(message *entity.Message) error
}
