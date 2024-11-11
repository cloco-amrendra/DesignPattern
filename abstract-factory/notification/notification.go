package notification

// EmailNotification interface for email notifications

type EmailNotification interface {
	SendEmail(message string) string
}

// SMSNotification interface for SMS notifications

type SMSNotification interface {
	SendSMS(message string) string
}
