package notification

// SMTPFactory is a factory for creating SMTP email and SMS notifications
type SMTPFactory struct{}

func (s *SMTPFactory) CreateEmailNotification() EmailNotification {
	return &SMTPEmailNotification{}
}

func (s *SMTPFactory) CreateSMSNotification() SMSNotification {
	return &SMTPSMSNotification{}
}
