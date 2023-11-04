package dto

import "github.com/mahdimehrabi/graph-interview/destination/internal/entity"

type MessageReq struct {
	Message    string `json:"message"`
	ReceivedAt int    `json:"ReceivedAt"`
}

func (m MessageReq) ToModel() *entity.Message {
	return &entity.Message{
		Message:    m.Message,
		ReceivedAt: m.ReceivedAt,
	}
}
