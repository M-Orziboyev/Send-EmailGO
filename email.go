package email

import (
	"bytes"
	"fmt"
	"html/template"
	"net/smtp"
)

type SendEmailRequest struct {
	To      []string
	Type    string
	Body    map[string]string
	Subject string
}

const (
	VerificationEmail   = "orziboyevmuzaffar@gmail.com"
	ForgotPasswordEmail = "forgot_password_email"
)

type Smtp struct {
	Sender   string
	Password string
}

func SendEmail(cfg *Smtp, req *SendEmailRequest) error {
	from := cfg.Sender
	to := req.To

	password := cfg.Password

	var body bytes.Buffer

	templatePath := "./html/index.html" //your template path (yo'lak)
	t, err := template.ParseFiles(templatePath)
	if err != nil {
		return err
	}

	err = t.Execute(&body, req.Body)
	if err!= nil {
        return err
    }

	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n" // must be
	subject := fmt.Sprintf("Subject: %s\n", req.Subject)
	msg := []byte(subject + mime + body.String())

	auth := smtp.PlainAuth("", from, password, "smtp.gmail.com")
	err = smtp.SendMail("smtp.gmail.com:587", auth, from, to, msg)
	if err != nil {
		return err
	}

	return nil
}
