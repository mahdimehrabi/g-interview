package internal

import (
	"github.com/mahdimehrabi/graph-interview/receiver/external/broker"
	infrastructures "github.com/mahdimehrabi/graph-interview/receiver/internal/infrastructure"
	"github.com/mahdimehrabi/graph-interview/receiver/internal/repository/message"
)

// this helps increase performance/scalability/reliability because requests will send more parallel and requests
// won't send and relies on only one socket
const socketConnectionCount = 100

// DPI singleton dependency injection
// for instances that use limited resources
var DPI *dpi

type dpi struct {
	BrokerSockets    []*broker.Socket
	BrokerRepository message.Message // broker repository
}

func SetupDPI(env *infrastructures.Env) {
	brokerSockets := make([]*broker.Socket, socketConnectionCount)
	for i := 0; i < socketConnectionCount; i++ {
		brokerSockets[i] = broker.NewSocket(env.BrokerAddress)
	}
	DPI = &dpi{
		BrokerSockets:    brokerSockets, // maybe our service use this sockets for other purposes than sending msgs to broker
		BrokerRepository: message.NewBroker(brokerSockets),
	}
}
