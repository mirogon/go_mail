package mail

import (
	"errors"
	"net/smtp"

	util "github.com/mirogon/go_util"
)

type GmailEmailSender struct {
	From string
	Pw   string
}

func CreateGmailEmailSender(from string, pw string) GmailEmailSender {
	return GmailEmailSender{From: from, Pw: pw}
}

func (emailSender GmailEmailSender) SendEmail(receiver string, subject string, message string) error {
	if !util.HasInternet() {
		return errors.New("No internet!")
	}
	from := emailSender.From
	password := emailSender.Pw
	to := []string{
		receiver,
	}

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	auth := smtp.PlainAuth("", from, password, smtpHost)

	msg := "Subject: " + subject + "\r\n" + "Content-Type: text/html; charset=UTF-8" + "\r\n\r\n" + message
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, []byte(msg))
	if err != nil {
		return err
	}
	return nil
}

func CreateEmailSender() GmailEmailSender {
	return GmailEmailSender{}
}
