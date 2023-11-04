package message

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	brokerSocket "github.com/mahdimehrabi/graph-interview/receiver/external/broker"
	"github.com/mahdimehrabi/graph-interview/receiver/external/utils"
	"github.com/mahdimehrabi/graph-interview/receiver/internal/entity"
)

const (
	workerCount          = 10
	saveMessageMethod    = "save_message"
	saveDeadlineDuration = 5 * time.Second
)

var ErrResourceNotAvailable = errors.New("resource is not available")

type broker struct {
	queue   chan *entity.Message
	sockets []*brokerSocket.Socket
}

func NewBroker(sockets []*brokerSocket.Socket) Message {
	b := &broker{
		sockets: sockets,
		queue:   make(chan *entity.Message, 10000),
	}
	go b.SaveQueue()
	return b
}

func (b broker) SaveQueue() {
	for i := 0; i < workerCount; i++ {
		go b.savingWorker()
	}
}

func (b broker) savingWorker() {
	for {
		msg := <-b.queue
		id := uuid.New().String()
		socket := b.sockets[utils.RandomNumber(len(b.sockets)-1)]

		deadline := time.NewTicker(saveDeadlineDuration)
		done := make(chan bool)
		go func(ch chan bool) {
			if _, err := socket.SendWaitJSON(msg, saveMessageMethod, id); err != nil {
				fmt.Printf("failed to save message %s trying again,err:%s", id, err.Error())
				time.Sleep(time.Millisecond * 100) // socket resend cool down
				return
			}
			done <- true
		}(done)
		select {
		case <-done:
			fmt.Printf("message %s saved succesfuly🥳 \n", msg.Message)
		case <-deadline.C: //deadline exceeded
			time.Sleep(100 * time.Millisecond)
			b.queue <- msg
		}
	}
}

func (b broker) Save(msg *entity.Message) error {
	if len(b.queue) >= 10000 {
		return ErrResourceNotAvailable
	}
	b.queue <- msg
	return nil
}
