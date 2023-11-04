package repository

import (
	"github.com/mahdimehrabi/graph-interview/destination/internal/entity"
	"github.com/stretchr/testify/mock"
)

// MockMessage is a mock implementation of the Message interface
type MockMessage struct {
	mock.Mock
}

// Save is a mock implementation for the Save method
func (m *MockMessage) Save(message *entity.Message) error {
	args := m.Called(message)
	return args.Error(0)
}
