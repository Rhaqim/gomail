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

	Newsletter(g)
}

func Newsletter(mail gomail.Gomail) {
	type Articles struct {
		Title   string
		Content string
	}

	email := &gomail.Email{
		Recipients:       []string{"recipient1@gmail.com", "recipient2@gmail.com"},
		Subject:          "Weekly Newsletter",
		Body:             "",
		TemplateFileName: "newsletter.html",
		Data: map[string]interface{}{
			"Articles": []Articles{
				{
					Title:   "Article 1",
					Content: "This is the content of article 1",
				},
				{
					Title:   "Article 2",
					Content: "This is the content of article 2",
				},
			},
		},
	}

	err := mail.SendEmail(email)
	if err != nil {
		log.Fatal(err)
	}
}
