package notification

// SESFactory is a factory for creating SES email and SMS notifications
type SESFactory struct{}

func (s *SESFactory) CreateEmailNotification() EmailNotification {
	return &SESEmailNotification{}
}

func (s *SESFactory) CreateSMSNotification() SMSNotification {
	return &SESSMSNotification{}
}
