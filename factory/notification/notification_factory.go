package notification

// NotificationServiceFactory is the factory interface for creating notification services
type NotificationServiceFactory interface {
	CreateNotificationService(notificationType string) NotificationService
}

// ConcreteNotificationServiceFactory implements the NotificationServiceFactory interface
type ConcreteNotificationServiceFactory struct{}

// CreateNotificationService creates and returns the appropriate NotificationService based on notificationType
func (f *ConcreteNotificationServiceFactory) CreateNotificationService(notificationType string) NotificationService {
	if notificationType == "sms" {
		return &SMSNotification{}
	}
	// Default to email notification service
	return &EmailNotification{}
}
