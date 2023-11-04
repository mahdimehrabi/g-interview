package message

import "github.com/mahdimehrabi/graph-interview/destination/internal/entity"

type Message interface {
	SaveAnalyze(message *entity.Message) error
}
