package message

import (
	"github.com/mahdimehrabi/graph-interview/receiver/application/http/response"
	"net/http"
)

type Message struct{}

func NewMessage() *Message {
	return &Message{}
}

func (g Message) Send(w http.ResponseWriter, _ *http.Request) {
	response.GenResponse(w, map[string]string{"ping": "pong🏓"}, "")
}
