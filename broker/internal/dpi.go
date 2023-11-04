package internal

import (
	"github.com/mahdimehrabi/graph-interview/broker/external/destination"
	infrastructures "github.com/mahdimehrabi/graph-interview/broker/internal/infrastructure"
	"github.com/mahdimehrabi/graph-interview/broker/internal/repository/message"
	destination2 "github.com/mahdimehrabi/graph-interview/broker/internal/repository/message/destination"
)

// this helps increase performance/scalability/reliability because requests will send more parallel and requests
// won't send and relies on only one socket
const socketConnectionCount = 100

// DPI singleton dependency injection
// for instances that use limited resources
var DPI *dpi

type dpi struct {
	BrokerSockets    []*destination.Socket
	BrokerRepository message.Message // destination repository
}

func SetupDPI(env *infrastructures.Env) {
	brokerSockets := make([]*destination.Socket, socketConnectionCount)
	for i := 0; i < socketConnectionCount; i++ {
		brokerSockets[i] = destination.NewSocket(env.DestinationAddress)
	}
	DPI = &dpi{
		BrokerSockets:    brokerSockets, // maybe our service use this sockets for other purposes than sending msgs to destination
		BrokerRepository: destination2.NewDestination(brokerSockets),
	}
}
