package main

import (
	"gomail"
	"log"
)

func main() {
	auth := gomail.EmailAuthConfig{
		Host:     "",
		Port:     0,
		Username: "",
		Password: "",
		From:     "",
	}

	var g gomail.Gomail = gomail.NewGoemail(auth, "templates")

	App(g)

}

func App(mail gomail.Gomail) {
	err := mail.SendEmail([]string{"me", "you"}, "Hello", "hello.html", "me")
	if err != nil {
		log.Fatal(err)
	}
}
