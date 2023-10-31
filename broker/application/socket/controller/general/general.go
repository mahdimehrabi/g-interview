package general

import (
	"github.com/mahdimehrabi/graph-interview/receiver/application/socket/response"
	"net"
)

type General struct{}

func NewGeneral() *General {
	return &General{}
}

func (g General) Ping(conn net.Conn) {
	response.SuccessResponse(conn, map[string]string{"ping": "pong🏓"}, "pong")
}

func (g General) NotDefined(conn net.Conn) {
	response.SuccessResponse(conn, nil, "method not defined")
}
