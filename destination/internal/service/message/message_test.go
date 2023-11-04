package message

import (
	"errors"
	"testing"

	"github.com/mahdimehrabi/graph-interview/destination/internal/entity"
	"github.com/mahdimehrabi/graph-interview/destination/mocks/repository"
	"github.com/stretchr/testify/assert"
)

func TestMessage_Save(t *testing.T) {
	mockMessage := new(repository.MockMessage)

	messageInstance := Message{messageRepo: mockMessage}

	sampleMessage := entity.NewMessage("sample message!")

	mockMessage.On("Save", sampleMessage).Return(nil)

	err := messageInstance.Save(sampleMessage)

	assert.Nil(t, err)

	mockMessage.AssertCalled(t, "Save", sampleMessage)
}

func TestMessage_Save_RepoError(t *testing.T) {
	mockMessage := new(repository.MockMessage)

	messageInstance := Message{messageRepo: mockMessage}

	sampleMessage := entity.NewMessage("sample message!")

	mockMessage.On("Save", sampleMessage).Return(errors.New("error"))

	err := messageInstance.Save(sampleMessage)

	assert.ErrorIs(t, err, ErrMessageInternal)

	mockMessage.AssertCalled(t, "Save", sampleMessage)
}
