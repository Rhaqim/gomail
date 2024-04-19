package gomail

import (
	"bytes"
	"fmt"
	"net/smtp"
	"sync"
	"text/template"

	"github.com/Rhaqim/gomail/errors"
)

// Goemail is an interface for sending emails and parsing templates.
type Gomail interface {
	// validate validates the email data.
	validate(kind ValidateKind) error

	// authenticate authenticates the email client.
	authenticate() (smtp.Auth, error)

	// parseTemplate parses the email template.
	parseTemplate() error

	// message constructs the email message.
	message() []byte

	// SendEmail sends an email to the specified recipients with the given subject, template, and data.
	SendEmail(mail *Email) error
}

func NewGomail(auth EmailAuthConfig, templateDir string) Gomail {

	return &GoemailConfig{
		Config:      auth,
		TemplateDir: templateDir,
	}
}

func (g *GoemailConfig) validate(kind ValidateKind) error {

	switch kind {
	case auth:
		if g.Config.Host == "" {
			return errors.ErrEmptyHost
		}

		if g.Config.Port == 0 {
			return errors.ErrEmptyPort
		}

		if g.Config.Username == "" {
			return errors.ErrEmptyUsername
		}

		if g.Config.Password == "" {
			return errors.ErrEmptyPassword
		}

		if g.Config.From == "" {
			return errors.ErrEmptyFrom
		}

		if g.TemplateDir == "" {
			return errors.ErrEmptyTemplateDir
		}
	case email:
		if len(g.email.Recipients) == 0 {
			return errors.ErrEmptyTo
		}
	}

	return nil
}

func (g *GoemailConfig) authenticate() (smtp.Auth, error) {

	err := g.validate(auth)
	if err != nil {
		return nil, err
	}

	auth := smtp.PlainAuth("", g.Config.Username, g.Config.Password, g.Config.Host)

	return auth, nil
}

func (g *GoemailConfig) parseTemplate() error {
	var data any
	var err error
	buf := new(bytes.Buffer)

	templFileName := g.TemplateDir + "/" + g.email.TemplateFileName

	t, err := template.ParseFiles(templFileName)
	if err != nil {
		return err
	}

	if g.email.Data != nil {
		data = g.email.Data
	} else {
		data = struct {
			Title string
			Body  string
		}{
			Title: g.email.Subject,
			Body:  g.email.Body,
		}
	}

	if err = t.Execute(buf, data); err != nil {
		return err
	}

	g.email.Body = buf.String()

	return nil
}

func (g *GoemailConfig) message() []byte {
	subject := "Subject: " + g.email.Subject + "!\n"

	// Set the "Content-Type" header to "text/html".
	header := make(map[string]string)
	header["Content-Type"] = "text/html; charset=\"UTF-8\";"
	header["MIME-version"] = "1.0;"

	var message []byte = []byte("From: " + g.Config.From + "\r\n")

	// Add the recipients to the message.
	var wg sync.WaitGroup
	wg.Add(len(g.email.Recipients))

	for _, recipient := range g.email.Recipients {
		go func(recipient string) {
			defer wg.Done()
			g.mutex.Lock()
			defer g.mutex.Unlock()
			message = append(message, []byte("To: "+recipient+"\r\n")...)
		}(recipient)
	}

	wg.Wait()

	message = append(message, []byte(subject+"\r\n")...)
	message = append(message, []byte(g.email.Body)...)

	// Add the headers to the message.
	for k, v := range header {
		message = append([]byte(k+": "+v+"\r\n"), message...)
	}

	return message
}

func (g *GoemailConfig) SendEmail(mail *Email) error {
	var err error

	g.email = mail

	err = g.validate(email)
	if err != nil {
		return err
	}

	err = g.parseTemplate()
	if err != nil {
		return err
	}

	auth, err := g.authenticate()
	if err != nil {
		return err
	}

	err = smtp.SendMail(g.Config.Host+":"+fmt.Sprint(g.Config.Port), auth, g.Config.Username, g.email.Recipients, g.message())
	if err != nil {
		return err
	}

	return err
}
