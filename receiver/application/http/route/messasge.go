package route

import (
	"github.com/mahdimehrabi/graph-interview/receiver/application/http/controller/message"
	"net/http"
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
	http.HandleFunc("/api/send", g.messageController.Send)
}
