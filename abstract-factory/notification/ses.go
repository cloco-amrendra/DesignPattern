package notification

// SESEmailNotification is a concrete implementation of EmailNotification
type SESEmailNotification struct{}

func (s *SESEmailNotification) SendEmail(message string) string {
	return "SES Email Sent: " + message
}

// SESSMSNotification is a concrete implementation of SMSNotification
type SESSMSNotification struct{}

func (s *SESSMSNotification) SendSMS(message string) string {
	return "SES SMS Sent: " + message
}
