package internal

import (
	"github.com/mahdimehrabi/graph-interview/receiver/external/broker"
	infrastructures "github.com/mahdimehrabi/graph-interview/receiver/internal/infrastructure"
	"github.com/mahdimehrabi/graph-interview/receiver/internal/repository/message"
)

// singleton dependency injection
// for instances that use limited resources
var DPI *dpi

type dpi struct {
	BrokerSocket     *broker.Socket
	BrokerRepository message.Message //broker repository
}

func SetupDPI(env *infrastructures.Env) {
	bs := broker.NewSocket(env.BrokerAddress)
	DPI = &dpi{
		BrokerSocket:     bs,
		BrokerRepository: message.NewBroker(bs),
	}
}
