package mail

type EmailSender interface {
	SendEmail(receiver string, subject string, message string, from string) error
}
