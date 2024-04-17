package main

import (
	"gomail"
	"log"
)

func main() {
	auth := gomail.EmailAuthConfig{
		Host:     "smtp.gmail.com",
		Port:     587,
		Username: "user",
		Password: "password",
		From:     "me@gmail.com",
	}

	var g gomail.Gomail = gomail.NewGoemail(auth, "example/templates")

	App(g)

}

func App(mail gomail.Gomail) {
	err := mail.SendEmail([]string{"me", "you"}, "hello.html", "Hello", "me")
	if err != nil {
		log.Fatal(err)
	}
}
