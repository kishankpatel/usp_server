package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMessageService(t *testing.T) {
	messageService := NewMessageService()

	t.Run("should return the decrypted text", func(t *testing.T) {
		key, msg := "PDQ3iVJB7ZZuv05aZO435g==", "Yj1x27hy-KeezEM6LGo4do093vSKiQZzJ7J5OLBLgojkireIrlQr"
		response, err := messageService.ReceiveMessage(key, msg)

		assert.Nil(t, err)
		assert.Equal(t, "Kishan Patel", response.DecryptedText)
	})

	t.Run("should return for arguments for decryption", func(t *testing.T) {
		key, msg := "SomeRandomKey", "SomeRandomMessage"
		response, err := messageService.ReceiveMessage(key, msg)

		assert.Nil(t, response)
		assert.NotEmpty(t, err.Error())
	})

	t.Run("should return the encrypted text", func(t *testing.T) {
		key, msg := "PDQ3iVJB7ZZuv05aZO435g==", "Kishan Patel"
		response, err := messageService.EncryptMessage(key, msg)

		assert.Nil(t, err)
		assert.NotEmpty(t, response)
	})

	t.Run("should return for arguments for decryption", func(t *testing.T) {
		key, msg := "SomeRandomKey", "SomeRandomEncryptedMessage"
		response, err := messageService.EncryptMessage(key, msg)

		assert.Nil(t, response)
		assert.NotEmpty(t, err.Error())
	})
}
