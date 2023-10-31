package message

import (
	"github.com/mahdimehrabi/graph-interview/receiver/application/socket/dto"
	"github.com/mahdimehrabi/graph-interview/receiver/application/socket/response"
	messageRepo "github.com/mahdimehrabi/graph-interview/receiver/internal/repository/message"
	"github.com/mahdimehrabi/graph-interview/receiver/internal/service/message"
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
		response.BadRequestErrorResponse(conn)
		return
	}
	msgEnt := data.ToModel()
	if err := g.messageService.Save(msgEnt); err != nil {
		response.InternalErrorResponse(conn)
		return
	}
	response.SuccessResponse(conn, nil, "message saved successfully")
}
