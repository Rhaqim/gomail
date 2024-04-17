# Gomail

Gomail is a Golang module that provides an abstraction for sending template emails to users in Golang applications. It allows users to provide their SMTP credentials and a folder for the email templates, and then sends the email using the specified template.

## Installation

To install Gomail, use the following command:

```bash
    go get github.com/rhaqim/gomail
```

## Usage

To use Gomail, you need to import the module and create a new instance of the `Gomail` struct. You can then use the `SendEmail` method to send an email.

```go
    package main

    import (
        "log"
        "github.com/rhaqim/gomail"
    )

    func main() {
        auth := gomail.EmailAuthConfig{
            Host:     "smtp.gmail.com",
            Port:     587,
            Username: "user",
            Password: "password",
            From:     "me@gmail.com",
        }

        templateDir := "templates"

        g := gomail.NewGoemail(auth, templatesDir)

        App(g)

    }

    func App(mail gomail.Gomail) {

        email := &gomail.Email{
            Recipients:       []string{"recipient1e@gmail.com", "recipiente2@gmail.com"},
            Subject:          "Hello",
            Body:             "Hello, this is a test email",
            TemplateFileName: "hello.html",
        }

        err := mail.SendEmail(email)
        if err != nil {
            log.Fatal(err)
        }
    }
```

## Contributing

To contribute to Gomail, fork the repository and create a new branch. Once you have made your changes, submit a pull request.
