package internal

import (
	"github.com/mahdimehrabi/graph-interview/destination/internal/repository/message"
	"github.com/mahdimehrabi/graph-interview/destination/internal/repository/message/stdout"
)

// singleton dependency injection
var DPI *dpi

type dpi struct {
	MessageRepo message.Message
}

func SetupDPI() {
	DPI = &dpi{
		MessageRepo: stdout.NewMessage(),
	}
}
