package stdout

import (
	"sync"
	"testing"

	"github.com/mahdimehrabi/graph-interview/destination/internal/entity"
	"github.com/stretchr/testify/assert"
)

func TestMessage_SaveAnalyze(t *testing.T) {
	// Create a message repository
	repo := message{
		lk: &sync.Mutex{},
	}

	// Create a sample message
	message := &entity.Message{
		Message: "Test Message",
	}

	// Call the SaveAnalyze method
	err := repo.SaveAnalyze(message)

	// Check for errors (in this case, there should be no error)
	assert.Nil(t, err)

	// Check if the Count and ByteSize are updated correctly
	assert.Equal(t, int64(1), repo.Count, "Count should be 1")
	size := int64(len([]byte(message.Message)))
	assert.Equal(t, size, repo.ByteSize, "ByteSize should be the length of the message")

	err = repo.SaveAnalyze(message)
	assert.Nil(t, err)

	// Check if the Count and ByteSize are updated correctly
	assert.Equal(t, int64(2), repo.Count, "Count should be 1")
	size += int64(len([]byte(message.Message)))
	assert.Equal(t, size, repo.ByteSize, "ByteSize should be the length of the message")
}
