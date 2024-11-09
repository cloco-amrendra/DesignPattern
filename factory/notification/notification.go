package notification

// NotificationService is the interface that all notification services implement
type NotificationService interface {
	SendNotification(message string) string
}

// EmailNotification is a concrete implementation of NotificationService for Email
type EmailNotification struct{}

func (e *EmailNotification) SendNotification(message string) string {
	return "Email Notification Sent: " + message
}

// SMSNotification is a concrete implementation of NotificationService for SMS
type SMSNotification struct{}

func (s *SMSNotification) SendNotification(message string) string {
	return "SMS Notification Sent: " + message
}
