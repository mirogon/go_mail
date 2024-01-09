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

func (emailSender GmailEmailSender) SendEmail(receiver string, subject string, message string, from_ string) error {
	if !util.HasInternet() {
		return errors.New("No internet!")
	}

	fromHeader := from_ + " <" + emailSender.From + ">"

	password := emailSender.Pw
	to := []string{
		receiver,
	}

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	auth := smtp.PlainAuth("", emailSender.From, password, smtpHost)

	msg := "Subject: " + subject + "\r\n" + "From: " + fromHeader + "\r\n" + "Content-Type: text/html; charset=UTF-8" + "\r\n\r\n" + message
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, emailSender.From, to, []byte(msg))
	if err != nil {
		return err
	}
	return nil
}

func CreateEmailSender() GmailEmailSender {
	return GmailEmailSender{}
}
