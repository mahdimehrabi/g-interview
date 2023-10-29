package entity

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestNewMessage(t *testing.T) {
	msg := "Hello, Testify!"

	// Create a new message
	message := NewMessage(msg)

	// Check if the message is not nil
	require.NotNil(t, message)

	// Check if the message content is set correctly
	require.Equal(t, msg, message.Message)

	// Check if ReceivedAt is not zero
	require.NotZero(t, message.ReceivedAt)

	// Check if ReceivedAt is not in the future
	require.LessOrEqual(t, message.ReceivedAt, time.Now().Unix())
}
