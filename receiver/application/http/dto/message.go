package dto

import "github.com/mahdimehrabi/graph-interview/receiver/internal/entity"

type MessageReq struct {
	Message string `json:"message"`
}

func (m MessageReq) ToModel() *entity.Message {
	return entity.NewMessage(m.Message)
}
