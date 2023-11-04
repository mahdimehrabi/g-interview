package message

import (
	"github.com/mahdimehrabi/graph-interview/broker/application/socket/dto"
	"github.com/mahdimehrabi/graph-interview/broker/application/socket/response"
	messageRepo "github.com/mahdimehrabi/graph-interview/broker/internal/repository/message"
	"github.com/mahdimehrabi/graph-interview/broker/internal/service/message"
	"net"
)

type Message struct {
	messageService *message.Message
}

func NewMessage() *Message {
	return &Message{
		message.NewMessage(messageRepo.NewBroker()),
	}
}

func (g Message) Save(conn net.Conn, req dto.Request) {
	data, ok := req.Data.(dto.MessageReq)
	if !ok {
		response.BadRequestErrorResponse(conn, req.ID)
		return
	}
	msgEnt := data.ToModel()
	if err := g.messageService.Save(msgEnt); err != nil {
		response.InternalErrorResponse(conn, req.ID)
		return
	}
	response.SuccessResponse(conn, nil, req.ID, "message saved successfully")
}
