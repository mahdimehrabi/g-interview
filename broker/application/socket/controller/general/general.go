package general

import (
	"net"

	"github.com/mahdimehrabi/graph-interview/broker/application/socket/dto"
	"github.com/mahdimehrabi/graph-interview/broker/application/socket/response"
)

type General struct{}

func NewGeneral() *General {
	return &General{}
}

func (g General) Ping(conn net.Conn, req dto.Request) {
	response.SuccessResponse(conn, map[string]string{"ping": "pong🏓"}, "pong", req.ID)
}

func (g General) NotDefined(conn net.Conn, req dto.Request) {
	response.BadRequestErrorResponse(conn, req.ID)
}
