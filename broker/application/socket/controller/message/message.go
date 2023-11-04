package message

import (
	"net"

	"github.com/mahdimehrabi/graph-interview/broker/application/socket/dto"
	"github.com/mahdimehrabi/graph-interview/broker/application/socket/response"
	"github.com/mahdimehrabi/graph-interview/broker/internal"
	"github.com/mahdimehrabi/graph-interview/broker/internal/service/message"
	"github.com/mitchellh/mapstructure"
)

type Message struct {
	messageService *message.Message
}

func NewMessage() *Message {
	return &Message{
		message.NewMessage(internal.DPI.BrokerRepository),
	}
}

func (g Message) Save(conn net.Conn, req dto.Request) {
	msgReq := dto.MessageReq{}
	if err := mapstructure.Decode(req.Data, &msgReq); err != nil {
		response.BadRequestErrorResponse(conn, req.ID)
		return
	}
	msgEnt := msgReq.ToModel()
	if err := g.messageService.Save(msgEnt); err != nil {
		response.InternalErrorResponse(conn, req.ID)
		return
	}
	response.SuccessResponse(conn, nil, req.ID, "message saved successfully")
}
