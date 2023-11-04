package stdout

import (
	"sync"

	"github.com/mahdimehrabi/graph-interview/destination/internal/entity"
	msgRepo "github.com/mahdimehrabi/graph-interview/destination/internal/repository/message"
	"github.com/rs/zerolog/log"
)

type message struct {
	Count    int64
	ByteSize int64
	lk       *sync.Mutex // s
}

func NewMessage() msgRepo.Message {
	return &message{Count: 0, ByteSize: 0, lk: &sync.Mutex{}}
}

func (r *message) SaveAnalyze(message *entity.Message) error {
	// because our + operations here are += we doesn't need for mutex and lock, but I use mutex anyway
	r.lk.Lock()
	r.ByteSize += int64(len([]byte(message.Message)))
	r.Count++
	r.lk.Unlock()
	log.Printf("message at timestamp %d saved 🔅\n", message.ReceivedAt)
	log.Printf("total size of messages %d total counts of messages %d ⚓\n", r.ByteSize, r.Count)
	return nil
}
