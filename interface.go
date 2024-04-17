package gomail

type EmailAuthConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	From     string
}

type Email struct {
	Recipients       []string
	Subject          string
	Body             string
	TemplateFileName string
}

type GoemailConfig struct {
	Config      EmailAuthConfig
	TemplateDir string
	Log         bool
	email       *Email
}

type ValidateKind string

const (
	auth  ValidateKind = "auth"
	email ValidateKind = "email"
)
