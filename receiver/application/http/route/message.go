package route

import (
	"net/http"

	"github.com/mahdimehrabi/graph-interview/receiver/application/http/controller/message"
)

type Message struct {
	messageController *message.Message
	mux               *http.ServeMux
}

func NewMessage(mux *http.ServeMux) *Message {
	return &Message{
		mux:               mux,
		messageController: message.NewMessage(),
	}
}

func (g Message) Handle() {
	g.mux.HandleFunc("/api/message/send", g.messageController.Send)
}
