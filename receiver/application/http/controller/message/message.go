package message

import (
	"encoding/json"
	"net/http"

	"github.com/mahdimehrabi/graph-interview/receiver/application/http/dto"
	"github.com/mahdimehrabi/graph-interview/receiver/application/http/response"
	"github.com/mahdimehrabi/graph-interview/receiver/internal"
	"github.com/mahdimehrabi/graph-interview/receiver/internal/service/message"
)

type Message struct {
	messageService *message.Message
}

func NewMessage() *Message {
	return &Message{
		message.NewMessage(internal.DPI.BrokerRepository),
	}
}

func (g Message) Send(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		response.MethodNotAllowedErrorResponse(w)
		return
	}
	req := dto.MessageReq{}
	d := json.NewDecoder(r.Body)
	if err := d.Decode(&req); err != nil {
		response.BadRequestErrorResponse(w)
		return
	}
	msgEnt := req.ToModel()
	if err := g.messageService.Save(msgEnt); err != nil {
		response.InternalErrorResponse(w)
		return
	}
	response.SuccessResponse(w, nil, "message sent successfully")
}
