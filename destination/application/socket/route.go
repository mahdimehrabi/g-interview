package socket

import (
	generalController "github.com/mahdimehrabi/graph-interview/destination/application/socket/controller/general"
	"github.com/mahdimehrabi/graph-interview/destination/application/socket/controller/message"
	"github.com/mahdimehrabi/graph-interview/destination/application/socket/dto"
	"net"
)

func HandleRoute(conn net.Conn, req dto.Request) {
	switch req.Method {
	case "save_message":
		msg := message.NewMessage() //transient dependency injection one instance per endpoint call to get best performance
		msg.Save(conn, req)
	case "ping":
		g := generalController.NewGeneral() //transient dependency injection one instance per endpoint call to get best performance
		g.Ping(conn, req)
	default:
		g := generalController.NewGeneral() //transient dependency injection one instance per endpoint call to get best performance
		g.NotDefined(conn, req)
	}
}
