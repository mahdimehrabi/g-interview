package message

import "github.com/mahdimehrabi/graph-interview/receiver/internal/entity"

type Message interface {
	Save(message *entity.Message) error
}
