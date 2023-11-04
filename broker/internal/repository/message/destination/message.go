package destination

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	brokerSocket "github.com/mahdimehrabi/graph-interview/broker/external/destination"
	"github.com/mahdimehrabi/graph-interview/broker/external/utils"
	"github.com/mahdimehrabi/graph-interview/broker/internal/entity"
	"github.com/mahdimehrabi/graph-interview/broker/internal/repository/message"
	"time"
)

const (
	workerCount       = 10
	saveMessageMethod = "save_message"
)

var ErrResourceNotAvailable = errors.New("resource is not available")

type destination struct {
	queue   chan *entity.Message
	sockets []*brokerSocket.Socket
}

func NewDestination(sockets []*brokerSocket.Socket) message.Message {
	b := &destination{
		sockets: sockets,
		queue:   make(chan *entity.Message, 10000),
	}
	go b.SaveQueue()
	return b
}

func (b destination) SaveQueue() {
	for i := 0; i < workerCount; i++ {
		go b.savingWorker()
	}
}

func (b destination) savingWorker() {
	for {
		msg := <-b.queue
		id := uuid.New().String()
		socket := b.sockets[utils.RandomNumber(len(b.sockets)-1)]

		if err := socket.SendWaitJSON(msg, saveMessageMethod, id); err != nil {
			fmt.Printf("failed to save message %s trying again,err:%s", id, err.Error())
			b.queue <- msg                     // add msg to end of the queue in case of error
			time.Sleep(time.Millisecond * 100) // socket resend cool down
			continue
		}
		fmt.Printf("message %s saved succesfuly😲 \n", id)
	}
}

func (b destination) Save(msg *entity.Message) error {
	if len(b.queue) >= 10000 {
		return ErrResourceNotAvailable
	}
	b.queue <- msg
	return nil
}
