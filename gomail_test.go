package gomail

import (
	"bytes"
	"testing"
)

func TestMessage(t *testing.T) {
	config := &GoemailConfig{
		email: &Email{
			Subject:    "Test Subject",
			Recipients: []string{"recipient1@example.com", "recipient2@example.com"},
			Body:       "Test Body",
		},
		Config: EmailAuthConfig{
			From: "sender@example.com",
		},
	}

	expectedMessage := []byte("MIME-version: 1.0;\r\n" +
		"Content-Type: text/html; charset=\"UTF-8\";\r\n" +
		"From: sender@example.com\r\n" +
		"To: recipient1@example.com\r\n" +
		"To: recipient2@example.com\r\n" +
		"Subject: Test Subject!\r\n" +
		"\r\n" +
		"Test Body")

	actualMessage := config.message()

	if !bytes.Equal(expectedMessage, actualMessage) {
		t.Errorf("Expected message:\n%s\n\nActual message:\n%s", expectedMessage, actualMessage)
	}
}
