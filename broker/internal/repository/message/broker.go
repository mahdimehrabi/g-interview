package message

import (
	extBroker "github.com/mahdimehrabi/graph-interview/broker/external/broker"
	"github.com/mahdimehrabi/graph-interview/broker/internal"
	"github.com/mahdimehrabi/graph-interview/broker/internal/entity"
)

type broker struct {
	socket *extBroker.Socket
}

func NewBroker() Message {
	return &broker{
		socket: internal.DPI.BrokerSocket,
	}
}

func (b broker) Save(msg *entity.Message) error {
	return b.socket.SendJSON(msg)
}
