package broker

import (
	extBroker "github.com/mahdimehrabi/graph-interview/receiver/external/broker"
	"github.com/mahdimehrabi/graph-interview/receiver/internal"
	"github.com/mahdimehrabi/graph-interview/receiver/internal/entity"
	"github.com/mahdimehrabi/graph-interview/receiver/internal/repository/message"
)

type broker struct {
	socket *extBroker.Socket
}

func NewBroker() message.Message {
	return &broker{
		socket: internal.DPI.BrokerSocket,
	}
}

func (b broker) Save(msg *entity.Message) error {

}
