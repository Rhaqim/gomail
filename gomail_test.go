package gomail

import (
	"bytes"
	"regexp"
	"sort"
	"strings"
	"testing"

	"github.com/Rhaqim/gomail/errors"
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

	// Sort expected and actual recipient lists
	sort.Strings(config.email.Recipients)
	actualRecipients := extractRecipients(actualMessage)
	sort.Strings(actualRecipients)

	if !bytes.Equal(expectedMessage, actualMessage) {
		t.Errorf("Expected message:\n%s\n\nActual message:\n%s", expectedMessage, actualMessage)
	}
}

func TestParseTemplate(t *testing.T) {
	config := &GoemailConfig{
		email: &Email{
			Subject:          "Test Subject",
			Recipients:       []string{"recipient1@example.com", "recipient2@example.com"},
			Body:             "Test Body",
			TemplateFileName: "hello.html",
			Data: map[string]interface{}{
				"Title": "Template Title",
				"Body":  "Template Body",
			},
		},
		TemplateDir: "example/templates",
	}

	err := config.parseTemplate()
	if err != nil {
		t.Errorf("Error parsing template: %v", err)
	}

	expectedBody := `<!DOCTYPE html>
	<html>
	<head>
		<title>My Go Template</title>
	</head>
	<body>
		<h1>Template Title</h1>
		<p>Template Body</p>
	</body>
	</html>`
	actualBody := config.email.Body

	// Normalize whitespace in expected and actual bodies
	expectedBodyNormalized := normalizeWhitespace(expectedBody)
	actualBodyNormalized := normalizeWhitespace(actualBody)

	if expectedBodyNormalized != actualBodyNormalized {
		t.Errorf("Expected body: %s, but got: %s", expectedBody, actualBody)
	}
}

func TestErrors(t *testing.T) {
	var tests = []struct {
		name     string
		err      error
		expected string
	}{
		{"ErrEmptyHost", errors.ErrEmptyHost, "email host is empty"},
		{"ErrEmptyPort", errors.ErrEmptyPort, "email port is empty"},
		{"ErrEmptyUsername", errors.ErrEmptyUsername, "email username is empty"},
		{"ErrEmptyPassword", errors.ErrEmptyPassword, "email password is empty"},
		{"ErrEmptyFrom", errors.ErrEmptyFrom, "email from is empty"},
		{"ErrEmptyTemplateDir", errors.ErrEmptyTemplateDir, "email template directory is empty"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.err.Error() != tt.expected {
				t.Errorf("Expected error: %s, but got: %s", tt.expected, tt.err.Error())
			}
		})
	}
}

// Helper function to normalize whitespace
func normalizeWhitespace(s string) string {
	s = strings.ReplaceAll(s, "\n", "")
	s = strings.ReplaceAll(s, "\t", "")
	s = strings.ReplaceAll(s, " ", "")
	return s
}

// Helper function to extract recipients from the actual message
func extractRecipients(message []byte) []string {
	// Define a regular expression to match "To:" lines in the message
	toRegex := regexp.MustCompile(`To: ([^\r\n]+)`)

	// Find all matches of the regular expression in the message
	matches := toRegex.FindAllSubmatch(message, -1)

	// Initialize a slice to store recipient addresses
	recipients := make([]string, len(matches))

	// Extract recipient addresses from the matched groups
	for i, match := range matches {
		if len(match) > 1 {
			recipients[i] = string(match[1])
		}
	}

	return recipients
}
