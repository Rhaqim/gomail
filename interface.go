package gomail

type EmailAuthConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	From     string
}

type Email struct {
	to               []string
	subject          string
	body             string
	templateFileName string
	data             interface{}
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
