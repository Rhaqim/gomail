package gomail

import "sync"

// EmailAuthConfig is a struct that contains the email authentication data.
// You can get the host, port, username and password from the email provider.
//
// Example:
//
//	EmailAuthConfig{
//		Host:     "smtp.gmail.com",
//		Port:     587,
//		Username: "user",
//		Password: "password",
//		From:     "me@gmail.com",
//	}
type EmailAuthConfig struct {
	Host     string // Host of the email provider
	Port     int    // Port of the email provider
	Username string // Username of the email provider
	Password string // Password of the email provider
	From     string // From address as the sender of the email
}

// Email is a struct that contains the email data.
//
// If data is provided, the template will be parsed with the data.
// Otherwise, the template will be parsed with the struct below.
// Ensure that the data keys match the template placeholders.
//
//	struct {
//		Title string // Title of the email
//		Body  string // Body of the email
//	}
type Email struct {
	Recipients       []string               // Recipients of the email
	Subject          string                 // Subject of the email
	Body             string                 // Body of the email
	TemplateFileName string                 // Name of the template file
	Data             map[string]interface{} // Data to be parsed with the template
}

type GoemailConfig struct {
	Config      EmailAuthConfig
	TemplateDir string
	Log         bool
	email       *Email
	mutex       sync.Mutex
}

type ValidateKind string

const (
	auth  ValidateKind = "auth"
	email ValidateKind = "email"
)
