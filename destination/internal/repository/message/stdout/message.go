package stdout

import (
	"fmt"
	"github.com/mahdimehrabi/graph-interview/destination/internal/entity"
	msgRepo "github.com/mahdimehrabi/graph-interview/destination/internal/repository/message"
	"sync"
)

type message struct {
	Count    int64
	ByteSize int64
	lk       *sync.Mutex //s
}

func NewMessage() msgRepo.Message {
	return &message{Count: 0, ByteSize: 0, lk: &sync.Mutex{}}
}

func (r *message) SaveAnalyze(message *entity.Message) error {
	//because our + operations here are += we doesn't need for mutex and lock, but I use mutex anyway
	r.lk.Lock()
	r.ByteSize += int64(len([]byte(message.Message)))
	r.Count++
	r.lk.Unlock()
	fmt.Printf("message %s saved 🔅\n", message.Message)
	fmt.Printf("total size of messages %d total counts of messages %d ⚓\n", r.ByteSize, r.Count)
	return nil
}
