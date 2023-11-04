package general

import (
	"github.com/mahdimehrabi/graph-interview/destination/application/socket/dto"
	"github.com/mahdimehrabi/graph-interview/destination/application/socket/response"
	"net"
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
