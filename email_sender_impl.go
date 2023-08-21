package mail

import (
	"fmt"
	"net/smtp"
)

type GmailEmailSender struct {
}

func (emailSender GmailEmailSender) SendEmail(receiver string, subject string, message string) error {
	from := "marcelroemersoftware@gmail.com"
	password := "aevmetmmhkprbukv"
	to := []string{
		receiver,
	}

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	auth := smtp.PlainAuth("", from, password, smtpHost)

	msg := "Subject: " + subject + "\r\n\r\n" + message
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, []byte(msg))
	if err != nil {
		return err
	}
	fmt.Println("Email sent!")
	return nil
}

func CreateEmailSender() GmailEmailSender {
	return GmailEmailSender{}
}
