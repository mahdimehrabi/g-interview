package message

import (
	"errors"
	"github.com/google/uuid"
	extBroker "github.com/mahdimehrabi/graph-interview/receiver/external/broker"
	"github.com/mahdimehrabi/graph-interview/receiver/internal/entity"
	"time"
)

const workerCount = 10
const saveMessageMethod = "save_message"

var ErrResourceNotAvailable = errors.New("resource is not available")

type broker struct {
	socket *extBroker.Socket
	queue  chan *entity.Message
}

func NewBroker(bs *extBroker.Socket) Message {
	b := &broker{
		socket: bs,
		queue:  make(chan *entity.Message, 10000),
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
		if err := b.socket.SendWaitJSON(msg, saveMessageMethod, uuid.New().String()); err != nil {
			b.queue <- msg                     //add msg to end of the queue in case of error
			time.Sleep(time.Millisecond * 100) //socket resend cool down
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
