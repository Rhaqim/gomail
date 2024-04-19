package main

import (
	"log"

	"github.com/Rhaqim/gomail"
)

func main() {
	auth := gomail.EmailAuthConfig{
		Host:     "smtp.gmail.com",
		Port:     587,
		Username: "user",
		Password: "password",
		From:     "me@gmail.com",
	}

	var g gomail.Gomail = gomail.NewGomail(auth, "example/templates")

	App(g)

}

func App(mail gomail.Gomail) {

	email := &gomail.Email{
		Recipients:       []string{"recipient1@gmail.com", "recipient2@gmail.com"},
		Subject:          "Hello",
		Body:             "Hello, this is a test email",
		TemplateFileName: "hello.html",
	}

	err := mail.SendEmail(email)
	if err != nil {
		log.Fatal(err)
	}
}
