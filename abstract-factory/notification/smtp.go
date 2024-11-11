package notification

// SMTPEmailNotification is a concrete implementation of EmailNotification

type SMTPEmailNotification struct{}

func (s *SMTPEmailNotification) SendEmail(message string) string {
	return "SMTP Email Sent: " + message
}

// SMTPSMSNotification is a concrete implementation of SMSNotification
type SMTPSMSNotification struct{}

func (s *SMTPSMSNotification) SendSMS(message string) string {
	return "SMTP SMS Sent: " + message
}
